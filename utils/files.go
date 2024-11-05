package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ericboy0224/patlytics-takehome/models"
)

func LoadPatents(filename string) ([]models.Patent, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.UseNumber() // To preserve number precision

	var patents []struct {
		ID                 int     `json:"id"`
		PublicationNumber  string  `json:"publication_number"`
		Title              string  `json:"title"`
		AISummary          string  `json:"ai_summary"`
		RawSourceURL       string  `json:"raw_source_url"`
		Assignee           string  `json:"assignee"`
		Inventors          string  `json:"inventors"`
		PriorityDate       string  `json:"priority_date"`
		ApplicationDate    string  `json:"application_date"`
		GrantDate          string  `json:"grant_date"`
		Abstract           string  `json:"abstract"`
		Description        string  `json:"description"`
		Claims             string  `json:"claims"`
		Jurisdictions      string  `json:"jurisdictions"`
		Classifications    string  `json:"classifications"`
		ApplicationEvents  string  `json:"application_events"`
		Citations          string  `json:"citations"`
		ImageURLs          string  `json:"image_urls"`
		Landscapes         string  `json:"landscapes"`
		CreatedAt          string  `json:"created_at"`
		UpdatedAt          string  `json:"updated_at"`
		PublishDate        string  `json:"publish_date"`
		CitationsNonPatent string  `json:"citations_non_patent"`
		Provenance         string  `json:"provenance"`
		AttachmentURLs     *string `json:"attachment_urls"`
	}

	if err := decoder.Decode(&patents); err != nil {
		return nil, err
	}

	// Convert to models.Patent
	result := make([]models.Patent, len(patents))
	for i, p := range patents {
		// Parse the time strings
		createdAt, err := time.Parse("2006-01-02 15:04:05.999999", p.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error parsing CreatedAt: %w", err)
		}
		updatedAt, err := time.Parse("2006-01-02 15:04:05.999999", p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error parsing UpdatedAt: %w", err)
		}

		result[i] = models.Patent{
			ID:                 p.ID,
			PublicationNumber:  p.PublicationNumber,
			Title:              p.Title,
			AISummary:          p.AISummary,
			RawSourceURL:       p.RawSourceURL,
			Assignee:           p.Assignee,
			Inventors:          p.Inventors,
			PriorityDate:       p.PriorityDate,
			ApplicationDate:    p.ApplicationDate,
			GrantDate:          p.GrantDate,
			Abstract:           p.Abstract,
			Description:        p.Description,
			Claims:             p.Claims,
			Jurisdictions:      p.Jurisdictions,
			Classifications:    p.Classifications,
			ApplicationEvents:  p.ApplicationEvents,
			Citations:          p.Citations,
			ImageURLs:          p.ImageURLs,
			Landscapes:         p.Landscapes,
			CreatedAt:          createdAt,
			UpdatedAt:          updatedAt,
			PublishDate:        p.PublishDate,
			CitationsNonPatent: p.CitationsNonPatent,
			Provenance:         p.Provenance,
			AttachmentURLs:     p.AttachmentURLs,
		}
	}

	return result, nil
}

func LoadCompanyProducts(filename string) ([]models.Company, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var wrapper struct {
		Companies []models.Company `json:"companies"`
	}

	if err := json.NewDecoder(file).Decode(&wrapper); err != nil {
		return nil, err
	}
	return wrapper.Companies, nil
}

func GetProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "..")
	return projectRoot
}
