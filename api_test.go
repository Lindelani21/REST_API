package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI(t *testing.T) {
	//  Initialize test DB (consider using a test database)
	InitDB()
	defer CloseDB()

	// Create the same router setup as in main()
	router := http.NewServeMux()
	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetBooksHandler(w, r)
		case http.MethodPost:
			CreateBookHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Create test server with the exact middleware stack from main()
	ts := httptest.NewServer(RecoveryMiddleware(LoggingMiddleware(router)))
	defer ts.Close()

	t.Run("Create and get book", func(t *testing.T) {
		// Test POST
		body := bytes.NewBufferString(`{"title":"Test Book","author":"Test Author"}`)
		resp, err := http.Post(ts.URL+"/books", "application/json", body)
		if err != nil {
			t.Fatalf("Failed to make POST request: %v", err)
		}
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status 201, got %d", resp.StatusCode)
		}

		// Test GET
		resp, err = http.Get(ts.URL + "/books")
		if err != nil {
			t.Fatalf("Failed to make GET request: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}
	})
}
