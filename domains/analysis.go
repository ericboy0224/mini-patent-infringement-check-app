package domains

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/ericboy0224/patlytics-takehome/models"
)

func AnalyzeInfringement(publicationNumber string, companyName string, patents []models.Patent, products []*models.ClientProduct) ([]*models.ClientProduct, error) {
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

	// Parse patent claims
	var patentClaims []string
	if err := json.Unmarshal([]byte(targetPatent.Claims), &patentClaims); err != nil {
		return nil, fmt.Errorf("failed to parse patent claims: %w", err)
	}

	// Filter products by company name and store with claim count
	type productWithClaims struct {
		product    *models.ClientProduct
		claimCount int
	}
	var infringingProducts []productWithClaims

	for _, product := range products {
		if product.CompanyName == companyName {
			commonClaims := findCommonClaims(patentClaims, product.Claims)
			if len(commonClaims) > 0 {
				infringingProducts = append(infringingProducts, productWithClaims{
					product:    product,
					claimCount: len(commonClaims),
				})
			}
		}
	}

	// Sort products by number of common claims (descending)
	sort.Slice(infringingProducts, func(i, j int) bool {
		return infringingProducts[i].claimCount > infringingProducts[j].claimCount
	})

	// Return the top 2 infringing products
	result := make([]*models.ClientProduct, 0, 2)
	for i := 0; i < len(infringingProducts) && i < 2; i++ {
		result = append(result, infringingProducts[i].product)
	}

	return result, nil
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
