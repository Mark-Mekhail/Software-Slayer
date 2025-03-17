package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ErrorResponse represents a standardized API error response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

/*
 * Decode decodes the request body into the given struct.
 * @param w: the response writer
 * @param r: the request
 * @param v: the struct to decode the request body into
 * @return error: an error if the decoding fails
 */
func Decode[T any](w http.ResponseWriter, r *http.Request, v *T) error {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return err
}

/*
 * RespondWithError sends a standardized JSON error response
 * @param w: the response writer
 * @param status: HTTP status code
 * @param message: error message
 */
func RespondWithError(w http.ResponseWriter, status int, message string) {
	log.Printf("ERROR: %s (Status: %d)", message, status)

	response := ErrorResponse{
		Status:  status,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// If encoding fails, fall back to a simple HTTP error
		http.Error(w, fmt.Sprintf("Internal server error: %s", err.Error()), http.StatusInternalServerError)
	}
}

/*
 * RespondWithJSON sends a standardized JSON response
 * @param w: the response writer
 * @param status: HTTP status code
 * @param payload: data to send in response
 */
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			// If encoding fails, respond with an error
			RespondWithError(w, http.StatusInternalServerError, "Failed to encode response")
		}
	}
}
