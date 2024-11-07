package main

import (
	"context"
	"log"

	"github.com/ericboy0224/patlytics-takehome/services"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := services.InitMongoDB(); err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}

	// Initialize collections with data
	ctx := context.Background()
	if err := services.InitializeCollections(ctx); err != nil {
		log.Fatalf("Failed to initialize collections: %v", err)
	}

	router := gin.Default()
	setupRoutes(router)
	log.Fatal(router.Run(":8080"))
}
