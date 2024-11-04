package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	setupRoutes(router)
	log.Fatal(router.Run(":8080")) // Start server on port 8080
}
