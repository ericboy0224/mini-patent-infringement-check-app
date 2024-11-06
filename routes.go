package main

import (
	"github.com/ericboy0224/patlytics-takehome/handlers"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	// Serve frontend static files
	router.Static("/", "./frontend/dist")

	// API routes
	v1 := router.Group("/patlytics/v1")
	{
		v1.POST("/infringement-check", handlers.HandleInfringementCheck)
	}

	// Fallback route for SPA
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})
}
