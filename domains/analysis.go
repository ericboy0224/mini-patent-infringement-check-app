package domains

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ericboy0224/patlytics-takehome/models"
	"github.com/firebase/genkit/go/plugins/dotprompt"
	"github.com/joho/godotenv"
	groq "github.com/jpoz/groq"
)

type InfringementAnalysis struct {
	ProductName            string   `json:"product_name"`
	InfringementLikelihood string   `json:"infringement_likelihood"`
	RelevantClaims         []int    `json:"relevant_claims"` // Changed to `int` based on provided data
	Explanation            string   `json:"explanation"`
	SpecificFeatures       []string `json:"specific_features"`
}

type AnalyzeResult struct {
	InfringingProducts    []InfringementAnalysis `json:"infringing_products"`
	OverallRiskAssessment string                 `json:"overall_risk_assessment"`
}

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: Error loading .env file: %v\n", err)
	}

	// Set the directory where your prompt files are located
	dotprompt.SetDirectory("prompts")
}

func PromptGenerator(patentClaims []string, products []models.Product) (string, error) {
	// Load the prompt template
	prompt, err := dotprompt.Open("patent_analysis")
	if err != nil {
		return "", fmt.Errorf("failed to load prompt template: %w", err)
	}

	// Clean and validate inputs
	cleanedClaims := make([]string, 0)
	for _, claim := range patentClaims {
		if trimmed := strings.TrimSpace(claim); trimmed != "" {
			cleanedClaims = append(cleanedClaims, trimmed)
		}
	}

	// Create product details array
	productDetails := make([]map[string]string, len(products))
	for i, product := range products {
		productDetails[i] = map[string]string{
			"name":        strings.TrimSpace(product.Name),
			"description": strings.TrimSpace(product.Description),
		}
	}

	// Prepare variables for the prompt
	variables := map[string]any{
		"patentClaims": strings.Join(cleanedClaims, "\n"),
		"products":     productDetails,
	}

	// Render the prompt
	renderedPrompt, err := prompt.RenderText(variables)
	if err != nil {
		return "", fmt.Errorf("failed to render prompt: %w", err)
	}

	return renderedPrompt, nil
}

func getGroqAPIKey() (string, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("GROQ_API_KEY environment variable is not set")
	}
	return apiKey, nil
}

func cleanResponseContent(content string) string {
	content = regexp.MustCompile("(?s)^```json\\n|\\n```$").ReplaceAllString(content, "")
	// Remove any newline escape sequences
	content = regexp.MustCompile(`\\n`).ReplaceAllString(content, "")
	return content
}

func AnalyzeInfringementWithGroq(patentClaims []string, products []models.Product) (*AnalyzeResult, error) {
	// Generate the prompt for all products
	renderedPrompt, err := PromptGenerator(patentClaims, products)
	if err != nil {
		return nil, err
	}

	// Get API key
	apiKey, err := getGroqAPIKey()
	if err != nil {
		return nil, err
	}
	groqClient := groq.NewClient(groq.WithAPIKey(apiKey))

	// Create chat completion request
	response, err := groqClient.CreateChatCompletion(groq.CompletionCreateParams{
		Model: "gemma-7b-it",
		Messages: []groq.Message{
			{
				Role:    "user",
				Content: renderedPrompt,
			},
		},
		Temperature: 0.3,
		MaxTokens:   8000, // Increased for multiple products
	})

	if err != nil {
		return nil, fmt.Errorf("failed to make Groq request: %w", err)
	}

	if len(response.Choices) == 0 {
		return nil, fmt.Errorf("no response from Groq")
	}

	content := response.Choices[0].Message.Content
	// Clean the response content
	cleanedContent := cleanResponseContent(content)

	var result AnalyzeResult
	if err := json.Unmarshal([]byte(cleanedContent), &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Validate the likelihood values
	for i := range result.InfringingProducts {
		switch result.InfringingProducts[i].InfringementLikelihood {
		case "High", "Moderate", "Low":
			// Valid values, do nothing
		default:
			result.InfringingProducts[i].InfringementLikelihood = "Low"
		}
	}

	return &result, nil
}
