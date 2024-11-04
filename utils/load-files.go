package utils

import (
	"encoding/json"
	"os"

	"github.com/ericboy0224/patlytics-takehome/models"
)

func LoadPatents(filename string) ([]models.Patent, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var patents []models.Patent
	if err := json.NewDecoder(file).Decode(&patents); err != nil {
		return nil, err
	}
	return patents, nil
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
