package handlers

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ericboy0224/patlytics-takehome/domains"
	"github.com/ericboy0224/patlytics-takehome/models"
	"github.com/ericboy0224/patlytics-takehome/utils"
	"github.com/gin-gonic/gin"
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

	// Load patents and products
	projectRoot := "."
	pathStr := filepath.Join(projectRoot, "data", "patents.json")
	patents, err := utils.LoadPatents(pathStr)
	if err != nil {
		c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to load patents: %v", err)))
		return
	}

	// Find target patent
	var targetPatent models.Patent
	for _, patent := range patents {
		if patent.PublicationNumber == request.PublicationNumber {
			targetPatent = patent
			break
		}
	}
	if targetPatent.PublicationNumber == "" {
		c.JSON(404, NewErrorResponse("Patent not found"))
		return
	}

	companies, err := utils.LoadCompanyProducts(filepath.Join(projectRoot, "data", "company_products.json"))
	if err != nil {
		c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to load company products: %v", err)))
		return
	}

	// Filter companies based on the requested company name
	var filteredCompanies []*models.Company
	for _, company := range companies {
		if strings.Contains(strings.ToUpper(company.Name), strings.ToUpper(request.CompanyName)) {
			filteredCompanies = append(filteredCompanies, &company)
		}
	}

	if len(filteredCompanies) == 0 {
		c.JSON(404, NewErrorResponse("Company not found"))
		return
	}

	claims, err := targetPatent.ExtractClaims()
	if err != nil {
		c.JSON(500, NewErrorResponse("Failed to extract claims: "+err.Error()))
		return
	}

	analysis, err := domains.AnalyzeInfringementWithGroq(claims, filteredCompanies[0].Products)
	if err != nil {
		c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to analyze products: %v", err)))
		return
	}

	var infringingProducts []map[string]interface{}
	for _, product := range analysis.InfringingProducts {
		infringingProducts = append(infringingProducts, map[string]interface{}{
			"product_name":            product.ProductName,
			"infringement_likelihood": product.InfringementLikelihood,
			"relevant_claims":         product.RelevantClaims,
			"explanation":             product.Explanation,
			"specific_features":       product.SpecificFeatures,
		})
	}

	responseData := gin.H{
		"patent_id":               request.PublicationNumber,
		"company_name":            filteredCompanies[0].Name,
		"infringing_products":     infringingProducts,
		"overall_risk_assessment": analysis.OverallRiskAssessment,
	}

	c.JSON(200, NewSuccessResponse(responseData, "Infringement check completed successfully"))
}

func HandleGetReports(c *gin.Context) {
	c.JSON(200, NewSuccessResponse("<reports here>", "Saved reports retrieved successfully"))
}
