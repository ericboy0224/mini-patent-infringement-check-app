package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpecificFeature string

type InfringingProduct struct {
	ProductName            string            `json:"product_name"`
	InfringementLikelihood string            `json:"infringement_likelihood"`
	RelevantClaims         []int             `json:"relevant_claims"`
	Explanation            string            `json:"explanation"`
	SpecificFeatures       []SpecificFeature `json:"specific_features"`
}

type PatentAnalysis struct {
	AnalysisID            string              `json:"analysis_id"`
	PatentID              string              `json:"patent_id"`
	CompanyName           string              `json:"company_name"`
	AnalysisDate          string              `json:"analysis_date"`
	TopInfringingProducts []InfringingProduct `json:"top_infringing_products"`
	OverallRiskAssessment string              `json:"overall_risk_assessment"`
}

type AnalysisRecord struct {
	ID                    primitive.ObjectID  `bson:"_id,omitempty" json:"analysis_id"`
	PatentID              string              `bson:"patent_id" json:"patent_id"`
	CompanyName           string              `bson:"company_name" json:"company_name"`
	AnalysisDate          time.Time           `bson:"analysis_date" json:"analysis_date"`
	InfringingProducts    []InfringingProduct `bson:"infringing_products" json:"infringing_products"`
	OverallRiskAssessment string              `bson:"overall_risk_assessment" json:"overall_risk_assessment"`
}
