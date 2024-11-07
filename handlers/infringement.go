package handlers

import (
	"fmt"
	"time"

	"github.com/ericboy0224/patlytics-takehome/domains"
	"github.com/ericboy0224/patlytics-takehome/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ericboy0224/patlytics-takehome/services"
)

func HandleInfringementCheck(c *gin.Context) {
	var request struct {
		PublicationNumber string `json:"patent_id"`
		CompanyName       string `json:"company_name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, NewErrorResponse("Invalid request format"))
		return
	}

	// Validate inputs
	if request.PublicationNumber == "" {
		c.JSON(400, NewErrorResponse("Publication number is required"))
		return
	}
	if request.CompanyName == "" {
		c.JSON(400, NewErrorResponse("Company name is required"))
		return
	}

	// Find target patent
	patentsCollection := services.GetPatentsCollection()
	var targetPatent models.Patent
	err := patentsCollection.FindOne(c, bson.M{"publication_number": request.PublicationNumber}).Decode(&targetPatent)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, NewErrorResponse("Patent not found"))
			return
		}
		c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to fetch patent: %v", err)))
		return
	}

	// Find company
	companiesCollection := services.GetCompaniesCollection()
	var company models.Company
	err = companiesCollection.FindOne(c, bson.M{
		"name": bson.M{"$regex": request.CompanyName, "$options": "i"},
	}).Decode(&company)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, NewErrorResponse("Company not found"))
			return
		}
		c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to fetch company: %v", err)))
		return
	}

	claims, err := targetPatent.ExtractClaims()
	if err != nil {
		c.JSON(500, NewErrorResponse("Failed to extract claims: "+err.Error()))
		return
	}

	// Check if analysis exists in MongoDB
	existingAnalysis, err := domains.GetExistingAnalysis(c, request.PublicationNumber, company.Name)
	if err != nil {
		c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to check existing analysis: %v", err)))
		return
	}

	var responseData gin.H
	if existingAnalysis != nil {
		responseData = gin.H{
			"analysis_id":             existingAnalysis.ID,
			"analysis_date":           existingAnalysis.AnalysisDate,
			"patent_id":               existingAnalysis.PatentID,
			"company_name":            existingAnalysis.CompanyName,
			"infringing_products":     existingAnalysis.InfringingProducts,
			"overall_risk_assessment": existingAnalysis.OverallRiskAssessment,
		}
	} else {
		analysis, err := domains.AnalyzeInfringementWithGroq(claims, company.Products)
		if err != nil {
			c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to analyze products: %v", err)))
			return
		}

		// Create new analysis record
		newAnalysis := &models.AnalysisRecord{
			PatentID:              request.PublicationNumber,
			CompanyName:           company.Name,
			AnalysisDate:          time.Now(),
			InfringingProducts:    analysis.InfringingProducts,
			OverallRiskAssessment: analysis.OverallRiskAssessment,
		}

		if err := domains.SaveAnalysis(c, newAnalysis); err != nil {
			c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to save analysis: %v", err)))
			return
		}

		responseData = gin.H{
			"analysis_id":             newAnalysis.ID,
			"analysis_date":           newAnalysis.AnalysisDate,
			"patent_id":               newAnalysis.PatentID,
			"company_name":            newAnalysis.CompanyName,
			"infringing_products":     newAnalysis.InfringingProducts,
			"overall_risk_assessment": newAnalysis.OverallRiskAssessment,
		}
	}

	c.JSON(200, NewSuccessResponse(responseData, "Infringement check completed successfully"))
}
