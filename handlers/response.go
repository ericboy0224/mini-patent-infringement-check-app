package handlers

// Response represents a standardized API response structure
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// NewSuccessResponse creates a success response
func NewSuccessResponse(data interface{}, message string) Response {
	return Response{
		Status:  "success",
		Data:    data,
		Message: message,
	}
}

// NewErrorResponse creates an error response
func NewErrorResponse(err string) Response {
	return Response{
		Status: "error",
		Error:  err,
	}
}
