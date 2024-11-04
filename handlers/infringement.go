package handlers

import (
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
	patents, err := utils.LoadPatents("data/patents.json")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to load patents"})
		return
	}

	companies, err := utils.LoadCompanyProducts("data/company_products.json")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to load company products"})
		return
	}

	// Transform products to client products
	var clientProducts []*models.ClientProduct
	for _, company := range companies {
		if company.Name == request.CompanyName {
			clientProducts, err = company.ToClientProducts()
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to transform products: " + err.Error()})
				return
			}
			break
		}
	}

	// Run infringement analysis
	infringingProducts, err := domains.AnalyzeInfringement(request.PublicationNumber, request.CompanyName, patents, clientProducts)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to analyze infringement: " + err.Error()})
		return
	}

	// Return detailed response
	c.JSON(200, gin.H{
		"status": "success",
		"data": gin.H{
			"infringing_products": infringingProducts,
		},
	})
}

func HandleGetReports(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Saved reports",
		"data":    "<reports here>",
	})
}
