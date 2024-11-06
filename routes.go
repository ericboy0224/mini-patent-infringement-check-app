package main

import (
	"github.com/ericboy0224/patlytics-takehome/handlers"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	router.POST("/infringement-check", handlers.HandleInfringementCheck)
}
