package models

import (
	"fmt"
	"strings"
)

// Company represents a company and its products
type Company struct {
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

// Product represents a product with its name and description
type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ClientProduct represents the client-side product structure
type ClientProduct struct {
	Name        string   `json:"product_name"`
	Summary     string   `json:"summary"`
	Claims      []string `json:"claims"`
	CompanyName string   `json:"company_name"`
}

// Companies is a wrapper for the array of companies
type products struct {
	Companies []Company `json:"companies"`
}

// ToClientProduct converts a Product to ClientProduct
func (p *Product) ToClientProduct(companyName string) (*ClientProduct, error) {
	// Split description into claims (you might want to adjust this logic)
	claims := strings.Split(p.Description, ". ")

	return &ClientProduct{
		Name:        p.Name,
		Summary:     p.Description,
		Claims:      claims,
		CompanyName: companyName,
	}, nil
}

// ToClientProducts converts all products in a company to ClientProducts
func (c *Company) ToClientProducts() ([]*ClientProduct, error) {
	clientProducts := make([]*ClientProduct, 0, len(c.Products))

	for _, product := range c.Products {
		clientProduct, err := product.ToClientProduct(c.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to convert product %s: %w", product.Name, err)
		}
		clientProducts = append(clientProducts, clientProduct)
	}

	return clientProducts, nil
}
