package services

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ericboy0224/patlytics-takehome/models"
	"github.com/ericboy0224/patlytics-takehome/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return fmt.Errorf("MONGODB_URI environment variable is not set")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	mongoClient = client
	return nil
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}

func GetAnalysisCollection() *mongo.Collection {
	return mongoClient.Database("patlytics").Collection("analyses")
}

func GetPatentsCollection() *mongo.Collection {
	return mongoClient.Database("patlytics").Collection("patents")
}

func GetCompaniesCollection() *mongo.Collection {
	return mongoClient.Database("patlytics").Collection("companies")
}

func InitializeCollections(ctx context.Context) error {
	// Load patents and products
	projectRoot := "."
	pathStr := filepath.Join(projectRoot, "data", "patents.json")
	rawPatents, err := utils.LoadPatents(pathStr)
	if err != nil {
		return fmt.Errorf("failed to load patents: %w", err)
	}

	// Convert patents to model structs
	patentDocs := make([]interface{}, len(rawPatents))
	for i, patent := range rawPatents {
		// Assuming patent is already a models.Patent since it's loaded through utils.LoadPatents
		patentDocs[i] = models.Patent{
			ID:                patent.ID,
			PublicationNumber: patent.PublicationNumber,
			Title:             patent.Title,
			AISummary:         patent.AISummary,
			RawSourceURL:      patent.RawSourceURL,
			Assignee:          patent.Assignee,
			Inventors:         patent.Inventors,
			Claims:            patent.Claims,
		}
	}

	// Insert patents
	patentsCollection := GetPatentsCollection()
	if _, err := patentsCollection.DeleteMany(ctx, bson.M{}); err != nil {
		return fmt.Errorf("failed to clear patents collection: %w", err)
	}
	if len(patentDocs) > 0 {
		if _, err := patentsCollection.InsertMany(ctx, patentDocs); err != nil {
			return fmt.Errorf("failed to insert patents: %w", err)
		}
	}

	// Load companies
	rawCompanies, err := utils.LoadCompanyProducts(filepath.Join(projectRoot, "data", "company_products.json"))
	if err != nil {
		return fmt.Errorf("failed to load companies: %w", err)
	}

	// Convert companies to model structs
	companyDocs := make([]interface{}, len(rawCompanies))
	for i, company := range rawCompanies {
		companyDocs[i] = models.Company{
			Name:     company.Name,
			Products: company.Products,
		}
	}

	// Insert companies
	companiesCollection := GetCompaniesCollection()
	if _, err := companiesCollection.DeleteMany(ctx, bson.M{}); err != nil {
		return fmt.Errorf("failed to clear companies collection: %w", err)
	}
	if len(companyDocs) > 0 {
		if _, err := companiesCollection.InsertMany(ctx, companyDocs); err != nil {
			return fmt.Errorf("failed to insert companies: %w", err)
		}
	}

	// Add this after existing collection initialization code
	counterCollection := mongoClient.Database("patlytics").Collection("counters")

	// Initialize the counter if it doesn't exist
	_, err = counterCollection.UpdateOne(
		ctx,
		bson.M{"_id": "analysisId"},
		bson.M{"$setOnInsert": bson.M{"seq": int64(0)}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return fmt.Errorf("failed to initialize counter: %w", err)
	}

	return nil
}

func GetNextAnalysisID(ctx context.Context) (int64, error) {
	counterCollection := mongoClient.Database("patlytics").Collection("counters")

	result := counterCollection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": "analysisId"},
		bson.M{"$inc": bson.M{"seq": int64(1)}},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	var counter struct {
		ID  string `bson:"_id"`
		Seq int64  `bson:"seq"`
	}

	if err := result.Decode(&counter); err != nil {
		return 0, fmt.Errorf("failed to get next ID: %w", err)
	}

	return counter.Seq, nil
}
