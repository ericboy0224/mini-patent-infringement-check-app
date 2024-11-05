package domains

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/ericboy0224/patlytics-takehome/models"
)

func AnalyzeInfringement(publicationNumber string, companyName string, patents []models.Patent, products []models.Product) (*models.PatentAnalysis, error) {
	// Input validation
	if len(patents) == 0 || len(products) == 0 {
		return nil, fmt.Errorf("empty patents or products list")
	}

	// Find the patent with the specified publication number
	var targetPatent models.Patent
	var found bool
	for _, patent := range patents {
		if patent.PublicationNumber == publicationNumber {
			targetPatent = patent
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("patent with publication number %s not found", publicationNumber)
	}

	// Extract claims from the target patent
	patentClaims, err := targetPatent.ExtractClaims()
	if err != nil {
		return nil, fmt.Errorf("failed to extract patent claims: %w", err)
	}

	// Filter products by company name and analyze potential infringement
	var infringingProducts []models.InfringingProduct

	for _, product := range products {
		// Create product claims from description
		productClaims := []string{product.Description}

		// Find common claims
		commonClaims := findCommonClaims(patentClaims, productClaims)

		if len(commonClaims) > 0 {
			infringingProducts = append(infringingProducts, models.InfringingProduct{
				ProductName:            product.Name,
				InfringementLikelihood: calculateLikelihood(len(commonClaims), len(patentClaims)),
				RelevantClaims:         commonClaims,
				Explanation:            generateDetailedExplanation(product.Name, commonClaims),
				SpecificFeatures:       extractFeatures(commonClaims),
			})
		}
	}

	// Sort products by number of matching claims (descending)
	sort.Slice(infringingProducts, func(i, j int) bool {
		return len(infringingProducts[i].RelevantClaims) > len(infringingProducts[j].RelevantClaims)
	})

	// Take top 2 products
	if len(infringingProducts) > 2 {
		infringingProducts = infringingProducts[:2]
	}

	// TODO use openAI for output
	// Create analysis result
	analysis := &models.PatentAnalysis{
		AnalysisID:            generateAnalysisID(),
		PatentID:              publicationNumber,
		CompanyName:           companyName,
		AnalysisDate:          time.Now().Format(time.RFC3339),
		TopInfringingProducts: infringingProducts,
		OverallRiskAssessment: calculateOverallRisk(infringingProducts),
	}

	return analysis, nil
}

func findCommonClaims(claims1, claims2 []string) []string {
	if len(claims1) > len(claims2) {
		claims1, claims2 = claims2, claims1 // Use smaller slice for map
	}

	claimSet := make(map[string]struct{}, len(claims1))
	for _, claim := range claims1 {
		claimSet[claim] = struct{}{}
	}

	common := make([]string, 0, len(claims1))
	for _, claim := range claims2 {
		if _, exists := claimSet[claim]; exists {
			common = append(common, claim)
		}
	}
	return common
}

// Helper functions
func calculateLikelihood(matchingClaims, totalClaims int) string {
	ratio := float64(matchingClaims) / float64(totalClaims)
	switch {
	case ratio > 0.7:
		return "High"
	case ratio > 0.3:
		return "Medium"
	default:
		return "Low"
	}
}

func extractFeatures(claims []string) []models.SpecificFeature {
	features := make([]models.SpecificFeature, len(claims))
	for i, claim := range claims {
		features[i] = models.SpecificFeature(claim)
	}
	return features
}

func generateAnalysisID() string {
	return fmt.Sprintf("ANL-%d", time.Now().UnixNano())
}

func calculateOverallRisk(products []models.InfringingProduct) string {
	if len(products) == 0 {
		return "Low"
	}

	highCount := 0
	for _, p := range products {
		if p.InfringementLikelihood == "High" {
			highCount++
		}
	}

	if highCount > 0 {
		return "High"
	}
	return "Medium"
}

func generateDetailedExplanation(productName string, claims []string) string {
	if len(claims) == 0 {
		return "No matching claims found"
	}

	return fmt.Sprintf(
		"Product '%s' potentially infringes on %d patent claims. The matching features include: %s",
		productName,
		len(claims),
		strings.Join(claims, "; "),
	)
}
