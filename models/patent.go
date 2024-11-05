package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Patent represents a patent record
type Patent struct {
	ID                 int       `json:"id"`
	PublicationNumber  string    `json:"publication_number"`
	Title              string    `json:"title"`
	AISummary          string    `json:"ai_summary"`
	RawSourceURL       string    `json:"raw_source_url"`
	Assignee           string    `json:"assignee"`
	Inventors          string    `json:"inventors"` // JSON string containing inventor objects
	PriorityDate       string    `json:"priority_date"`
	ApplicationDate    string    `json:"application_date"`
	GrantDate          string    `json:"grant_date"`
	Abstract           string    `json:"abstract"`
	Description        string    `json:"description"`
	Claims             string    `json:"claims"` // JSON string containing claim objects
	Jurisdictions      string    `json:"jurisdictions"`
	Classifications    string    `json:"classifications"` // JSON string containing classification objects
	ApplicationEvents  string    `json:"application_events"`
	Citations          string    `json:"citations"` // JSON string containing citation objects
	ImageURLs          string    `json:"image_urls"`
	Landscapes         string    `json:"landscapes"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	PublishDate        string    `json:"publish_date"`
	CitationsNonPatent string    `json:"citations_non_patent"`
	Provenance         string    `json:"provenance"`
	AttachmentURLs     *string   `json:"attachment_urls"` // Pointer since it can be null
}

// PatentDTO represents a simplified patent record for client-side use
type PatentDTO struct {
	ID     string   `json:"patent_id"`
	Claims []string `json:"claims"`
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
