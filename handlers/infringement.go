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
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	// Validate inputs
	if request.PublicationNumber == "" {
		c.JSON(400, gin.H{"error": "Publication number is required"})
		return
	}
	if request.CompanyName == "" {
		c.JSON(400, gin.H{"error": "Company name is required"})
		return
	}

	// Load patents and products
	projectRoot := "." // Use relative path since we're in the container
	pathStr := filepath.Join(projectRoot, "data", "patents.json")
	patents, err := utils.LoadPatents(pathStr)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to load patents: %v", err)})
		return
	}

	companies, err := utils.LoadCompanyProducts(filepath.Join(projectRoot, "data", "company_products.json"))
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to load company products: %v", err)})
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
		c.JSON(404, gin.H{"error": "Company not found"})
		return
	}

	// Run infringement analysis
	analysis, err := domains.AnalyzeInfringement(request.PublicationNumber, request.CompanyName, patents, filteredCompanies[0].Products)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to analyze infringement: " + err.Error()})
		return
	}

	// Return detailed response
	c.JSON(200, gin.H{
		"status": "success",
		"data":   analysis,
	})
}

func HandleGetReports(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Saved reports",
		"data":    "<reports here>",
	})
}
