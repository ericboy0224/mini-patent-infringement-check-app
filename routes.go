package main

import (
	"github.com/ericboy0224/patlytics-takehome/handlers"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	// API routes
	v1 := router.Group("/patlytics/v1")
	{
		v1.POST("/infringement-check", handlers.HandleInfringementCheck)
		v1.GET("/patents", handlers.HandleGetPatentList)
		v1.GET("/companies", handlers.HandleGetCompanyList)
	}

	// Serve frontend static files
	router.Static("/assets", "./frontend/dist/assets")

	// Fallback route for SPA
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})
}
