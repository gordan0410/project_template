package helper

import (
	"encoding/json"
	"net/http"
)

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// WriteJSONResponse writes a JSON response to the HTTP response writer
func WriteJSONResponse(w http.ResponseWriter, statusCode int, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// WriteSuccessResponse writes a successful JSON response
func WriteSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	WriteJSONResponse(w, statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// WriteErrorResponse writes an error JSON response
func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}
	
	WriteJSONResponse(w, statusCode, APIResponse{
		Success: false,
		Message: message,
		Error:   errorMsg,
	})
}