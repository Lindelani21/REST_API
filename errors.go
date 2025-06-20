package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func WriteError(w http.ResponseWriter, status int, err error, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   err.Error(),
		Message: message,
	})
}

func LogError(r *http.Request, err error) {
	log.Printf(
		"Error: %s | Method: %s | Path: %s",
		err.Error(),
		r.Method,
		r.URL.Path,
	)
}