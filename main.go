package main

import (
	"log"
	"net/http"
	"strings"
	"strconv"
	//"myapi/docs" 
	//"github.com/swaggo/http-swagger"
)

// @title Book API
// @version 1.0
// @description REST API for book management
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
func main() {
    InitDB()
    defer CloseDB()

    // Create a new ServeMux
    router := http.NewServeMux()

	router.HandleFunc("/books", booksCollectionHandler)
    router.HandleFunc("/books/", bookItemHandler)
	router.HandleFunc("/swagger/*", httpSwagger.WrapHandler)

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
    
    router.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
        idStr := strings.TrimPrefix(r.URL.Path, "/books/")
        if _, err := strconv.Atoi(idStr); err != nil {
            http.NotFound(w, r)
            return
        }
        
        switch r.Method {
        case http.MethodPut:
            UpdateBookHandler(w, r)
        case http.MethodDelete:
            DeleteBookHandler(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    stack := RecoveryMiddleware(LoggingMiddleware(router))

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", stack))
}


func booksCollectionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetBooksHandler(w, r)
	case http.MethodPost:
		CreateBookHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func bookItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		UpdateBookHandler(w, r)
	case http.MethodDelete:
		DeleteBookHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}