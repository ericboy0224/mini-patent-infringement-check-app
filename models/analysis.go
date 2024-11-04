package models

type SpecificFeature string

type InfringingProduct struct {
	ProductName            string            `json:"product_name"`
	InfringementLikelihood string            `json:"infringement_likelihood"`
	RelevantClaims         []string          `json:"relevant_claims"`
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
