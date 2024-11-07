package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Patent represents a patent record
type Patent struct {
	ID                 int       `json:"id" bson:"id"`
	PublicationNumber  string    `json:"publication_number" bson:"publication_number"`
	Title              string    `json:"title" bson:"title"`
	AISummary          string    `json:"ai_summary" bson:"ai_summary"`
	RawSourceURL       string    `json:"raw_source_url" bson:"raw_source_url"`
	Assignee           string    `json:"assignee" bson:"assignee"`
	Inventors          string    `json:"inventors" bson:"inventors"` // JSON string containing inventor objects
	PriorityDate       string    `json:"priority_date" bson:"priority_date"`
	ApplicationDate    string    `json:"application_date" bson:"application_date"`
	GrantDate          string    `json:"grant_date" bson:"grant_date"`
	Abstract           string    `json:"abstract" bson:"abstract"`
	Description        string    `json:"description" bson:"description"`
	Claims             string    `json:"claims" bson:"claims"` // JSON string containing claim objects
	Jurisdictions      string    `json:"jurisdictions" bson:"jurisdictions"`
	Classifications    string    `json:"classifications" bson:"classifications"` // JSON string containing classification objects
	ApplicationEvents  string    `json:"application_events" bson:"application_events"`
	Citations          string    `json:"citations" bson:"citations"` // JSON string containing citation objects
	ImageURLs          string    `json:"image_urls" bson:"image_urls"`
	Landscapes         string    `json:"landscapes" bson:"landscapes"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bson:"updated_at"`
	PublishDate        string    `json:"publish_date" bson:"publish_date"`
	CitationsNonPatent string    `json:"citations_non_patent" bson:"citations_non_patent"`
	Provenance         string    `json:"provenance" bson:"provenance"`
	AttachmentURLs     *string   `json:"attachment_urls" bson:"attachment_urls"` // Pointer since it can be null
}

// ClaimItem represents the structure of each claim in the JSON
type ClaimItem struct {
	Num  string `json:"num"`
	Text string `json:"text"`
}

// ExtractClaims parses the Claims JSON string and returns a slice of claim texts
func (p *Patent) ExtractClaims() ([]string, error) {
	// Parse the claims JSON string into array of ClaimItem
	var claimItems []ClaimItem
	if err := json.Unmarshal([]byte(p.Claims), &claimItems); err != nil {
		return nil, fmt.Errorf("failed to parse claims: %w", err)
	}

	// Extract just the text from each claim
	claims := make([]string, len(claimItems))
	for i, claim := range claimItems {
		claims[i] = claim.Text
	}

	return claims, nil
}
