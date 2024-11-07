package models

import (
	"time"
)

type SpecificFeature string

type InfringingProduct struct {
	ProductName            string            `json:"product_name"`
	InfringementLikelihood string            `json:"infringement_likelihood"`
	RelevantClaims         []int             `json:"relevant_claims"`
	Explanation            string            `json:"explanation"`
	SpecificFeatures       []SpecificFeature `json:"specific_features"`
}

type AnalysisRecord struct {
	ID                    int64               `bson:"_id" json:"analysis_id"`
	PatentID              string              `bson:"patent_id" json:"patent_id"`
	CompanyName           string              `bson:"company_name" json:"company_name"`
	AnalysisDate          time.Time           `bson:"analysis_date" json:"analysis_date"`
	InfringingProducts    []InfringingProduct `bson:"infringing_products" json:"infringing_products"`
	OverallRiskAssessment string              `bson:"overall_risk_assessment" json:"overall_risk_assessment"`
}
