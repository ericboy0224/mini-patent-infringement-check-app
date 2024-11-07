package handlers

import (
	"fmt"

	"github.com/ericboy0224/patlytics-takehome/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleGetPatentList(c *gin.Context) {
	collection := services.GetPatentsCollection()
	findOptions := options.Find().SetProjection(bson.D{{Key: "publication_number", Value: 1}})
	cursor, err := collection.Find(c, bson.M{}, findOptions)
	if err != nil {
		c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to fetch patents: %v", err)))
		return
	}
	defer cursor.Close(c)

	var patents []string
	for cursor.Next(c) {
		var patent struct {
			PublicationNumber string `bson:"publication_number"`
		}
		if err := cursor.Decode(&patent); err != nil {
			c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to decode patent: %v", err)))
			return
		}
		patents = append(patents, patent.PublicationNumber)
	}

	c.JSON(200, NewSuccessResponse(patents, "Patent list retrieved successfully"))
}

func HandleGetCompanyList(c *gin.Context) {
	collection := services.GetCompaniesCollection()
	findOptions := options.Find().SetProjection(bson.D{{Key: "name", Value: 1}})
	cursor, err := collection.Find(c, bson.M{}, findOptions)
	if err != nil {
		c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to fetch companies: %v", err)))
		return
	}
	defer cursor.Close(c)

	var companies []string
	for cursor.Next(c) {
		var company struct {
			Name string `bson:"name"`
		}
		if err := cursor.Decode(&company); err != nil {
			c.JSON(500, NewErrorResponse(fmt.Sprintf("Failed to decode company: %v", err)))
			return
		}
		companies = append(companies, company.Name)
	}

	c.JSON(200, NewSuccessResponse(companies, "Company list retrieved successfully"))
}
