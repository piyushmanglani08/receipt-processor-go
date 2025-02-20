package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteJSON encodes data as JSON and writes it to the response.
func WriteJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

// WriteError writes an error message as a JSON response with a given status code.
func WriteError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": message}); err != nil {
		log.Printf("Error writing error response: %v", err)
	}
}
