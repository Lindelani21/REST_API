package main

import (
	"encoding/json"
	"net/http"
	"context"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		WriteError(w, http.StatusBadRequest, err, "Invalid request")
		return
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err, "Failed to hash password")
		return
	}

	_, err = DB.Exec(
		context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		user.Username, hashedPassword,
	)
	
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err, "Failed to create user")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, err, "Invalid request")
		return
	}

	var user User
	err := DB.QueryRow(
		context.Background(),
		"SELECT id, username, password FROM users WHERE username = $1",
		req.Username,
	).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		WriteError(w, http.StatusUnauthorized, nil, "Invalid credentials")
		return
	}

	if !CheckPasswordHash(req.Password, user.Password) {
		WriteError(w, http.StatusUnauthorized, nil, "Invalid credentials")
		return
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err, "Failed to generate token")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}