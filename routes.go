package main

import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine) {
	router.POST("/infringement-check", handleInfringementCheck)
	router.GET("/saved-reports", handleGetReports) // Optional
}

func handleInfringementCheck(c *gin.Context) {
	var request struct {
		PatentID    string `json:"patent_id"`
		CompanyName string `json:"company_name"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	// Load data and run analysis
	c.JSON(200, gin.H{"message": "Infringement check completed", "data": "<results here>"})
}

func handleGetReports(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Saved reports", "data": "<reports here>"})
}
