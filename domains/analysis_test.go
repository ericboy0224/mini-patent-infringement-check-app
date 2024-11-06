package domains

import (
	"os"
	"testing"
)

func TestGetGroqAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		envValue    string
		wantErr     bool
		expectedErr string
	}{
		{
			name:        "valid API key",
			envValue:    "test-api-key",
			wantErr:     false,
			expectedErr: "",
		},
		{
			name:        "missing API key",
			envValue:    "",
			wantErr:     true,
			expectedErr: "GROQ_API_KEY environment variable is not set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variable
			if tt.envValue != "" {
				os.Setenv("GROQ_API_KEY", tt.envValue)
				defer os.Unsetenv("GROQ_API_KEY")
			} else {
				os.Unsetenv("GROQ_API_KEY")
			}

			// Test the function
			got, err := getGroqAPIKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("getGroqAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.expectedErr {
				t.Errorf("getGroqAPIKey() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if !tt.wantErr && got != tt.envValue {
				t.Errorf("getGroqAPIKey() = %v, want %v", got, tt.envValue)
			}
		})
	}
}

func TestCleanResponseContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "complex json with nested structure",
			input:    "```json\n{\n  \"infringing_products\": [\n    {\n      \"product_name\": \"Walmart Shopping App\",\n      \"infringement_likelihood\": \"High\",\n      \"relevant_claims\": [1, 7, 8, 14, 15, 16, 27, 28, 29, 33, 34, 35],\n      \"explanation\": \"The Walmart Shopping App directly integrates with the advertised products displayed on the mobile device, allowing for seamless addition of those products to the shopping list. This functionality aligns with claims 1, 7, 8, 14, 15, 16, 27, 28, 29, 33, 34, 35. Additionally, the app offers features like targeted advertisements, tracking payloads, and automatic list generation, aligning with claims 27, 28, 29.\"\n      \"specific_features\": [\"Integrated product selection from ads\", \"Automatic list generation\", \"Targeted advertisements\"]\n    },\n    {\n      \"product_name\": \"Walmart+ Membership\",\n      \"infringement_likelihood\": \"Moderate\",\n      \"relevant_claims\": [1, 7, 8, 14, 15, 16],\n      \"explanation\": \"Walmart+ offers smart shopping list synchronization, which aligns with claims 1, 7, 8, 14, 15, 16. This feature allows users to seamlessly transfer product selections from advertisements to their online shopping list through the app. However, the level of infringement is considered moderate due to the additional features offered by the Walmart Shopping App.\"\n      \"specific_features\": [\"Smart shopping list synchronization\"]\n    }\n  ],\n  \"overall_risk_assessment\": \"The top two products analyzed, Walmart Shopping App and Walmart+, both offer functionalities that infringe upon the claimed methods and systems. Both apps enable users to seamlessly add products from advertisements to their shopping lists, aligning with multiple claims. The risk of infringement is considered high for the Walmart Shopping App due to its direct integration with advertised products, while the risk for Walmart+ is moderate due to its additional features.\"\n}\n```",
			expected: "{\n  \"infringing_products\": [\n    {\n      \"product_name\": \"Walmart Shopping App\",\n      \"infringement_likelihood\": \"High\",\n      \"relevant_claims\": [1, 7, 8, 14, 15, 16, 27, 28, 29, 33, 34, 35],\n      \"explanation\": \"The Walmart Shopping App directly integrates with the advertised products displayed on the mobile device, allowing for seamless addition of those products to the shopping list. This functionality aligns with claims 1, 7, 8, 14, 15, 16, 27, 28, 29, 33, 34, 35. Additionally, the app offers features like targeted advertisements, tracking payloads, and automatic list generation, aligning with claims 27, 28, 29.\"\n      \"specific_features\": [\"Integrated product selection from ads\", \"Automatic list generation\", \"Targeted advertisements\"]\n    },\n    {\n      \"product_name\": \"Walmart+ Membership\",\n      \"infringement_likelihood\": \"Moderate\",\n      \"relevant_claims\": [1, 7, 8, 14, 15, 16],\n      \"explanation\": \"Walmart+ offers smart shopping list synchronization, which aligns with claims 1, 7, 8, 14, 15, 16. This feature allows users to seamlessly transfer product selections from advertisements to their online shopping list through the app. However, the level of infringement is considered moderate due to the additional features offered by the Walmart Shopping App.\"\n      \"specific_features\": [\"Smart shopping list synchronization\"]\n    }\n  ],\n  \"overall_risk_assessment\": \"The top two products analyzed, Walmart Shopping App and Walmart+, both offer functionalities that infringe upon the claimed methods and systems. Both apps enable users to seamlessly add products from advertisements to their shopping lists, aligning with multiple claims. The risk of infringement is considered high for the Walmart Shopping App due to its direct integration with advertised products, while the risk for Walmart+ is moderate due to its additional features.\"\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanResponseContent(tt.input)
			if result != tt.expected {
				t.Errorf("cleanResponseContent() = %v, want %v", result, tt.expected)
			}
		})
	}
}
