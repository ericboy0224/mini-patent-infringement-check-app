package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleInfringementCheckValidation(t *testing.T) {
	// Set up test environment
	gin.SetMode(gin.TestMode)

	// Create test router
	router := gin.New()
	router.POST("/infringement", HandleInfringementCheck)

	tests := []struct {
		name       string
		body       map[string]interface{}
		wantStatus int
		wantError  string
	}{
		{
			name:       "Missing patent_id",
			body:       map[string]interface{}{"company_name": "Test Company"},
			wantStatus: http.StatusBadRequest,
			wantError:  "Publication number is required",
		},
		{
			name:       "Missing company_name",
			body:       map[string]interface{}{"patent_id": "US123456"},
			wantStatus: http.StatusBadRequest,
			wantError:  "Company name is required",
		},
		{
			name:       "Invalid JSON",
			body:       nil,
			wantStatus: http.StatusBadRequest,
			wantError:  "Invalid request format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jsonBody []byte
			if tt.body != nil {
				jsonBody, _ = json.Marshal(tt.body)
			} else {
				jsonBody = []byte(`{invalid json}`)
			}

			req := httptest.NewRequest("POST", "/infringement", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)

			var response struct {
				Success bool   `json:"success"`
				Error   string `json:"error"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.False(t, response.Success)
			assert.Equal(t, tt.wantError, response.Error)
		})
	}
}
