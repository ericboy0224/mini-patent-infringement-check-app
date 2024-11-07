package models

// Company represents a company and its products
type Company struct {
	Name     string    `json:"name" bson:"name"`
	Products []Product `json:"products" bson:"products"`
}

// Product represents a product with its name and description
type Product struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

// ClientProduct represents the client-side product structure
type ClientProduct struct {
	Name        string   `json:"product_name" bson:"product_name"`
	Summary     string   `json:"summary" bson:"summary"`
	Claims      []string `json:"claims" bson:"claims"`
	CompanyName string   `json:"company_name" bson:"company_name"`
}

// Companies is a wrapper for the array of companies
type products struct {
	Companies []Company `json:"companies" bson:"companies"`
}
