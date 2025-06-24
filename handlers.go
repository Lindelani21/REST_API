package main

import (
	"encoding/json"
	"net/http"
	"context"
	"strconv"
    "strings"
	//"github.com/jackc/pgx/v5"
)

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(context.Background(), "SELECT id, title , author FROM books")
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author)
		if err != nil {
			http.Error(w, "Failed to scan books", http.StatusInternalServerError)
			return
		}
		books = append(books, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// POST /books - Add a new book
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
    var book Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        WriteError(w, http.StatusBadRequest, err, "Invalid request format")
        return
    }

    if book.Title == "" || book.Author == "" {
        WriteError(w, http.StatusBadRequest, 
            nil, "Title and author are required")
        return
    }

    // Fixed: Properly declare err variable
    _, err := DB.Exec(
        context.Background(),
        "INSERT INTO books (title, author) VALUES ($1, $2)",
        book.Title, book.Author,
    )

    if err != nil {
        WriteError(w, http.StatusInternalServerError, err, "Failed to create book")
        LogError(r, err)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Book created!"})
}

// PUT /books/:id - Update a book
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Decode request body
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update in database
	_, err = DB.Exec(
		context.Background(),
		"UPDATE books SET title=$1, author=$2 WHERE id=$3",
		book.Title, book.Author, id,
	)

	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book updated"})
}

// DELETE /books/:id - Delete a book
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL
	idStr := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	// Delete from database
	_, err = DB.Exec(context.Background(), "DELETE FROM books WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
}
