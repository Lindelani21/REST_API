package main

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Printf(
				"Method: %s | Path: %s | Duration: %v",
				r.Method,
				r.URL.Path,
				time.Since(start),
			)
		}()
		next.ServeHTTP(w, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				WriteError(w, http.StatusInternalServerError, 
					nil, "Internal server error")
				LogError(r, err.(error))
			}
		}()
		next.ServeHTTP(w, r)
	})
}