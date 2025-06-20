package main

import (
	"context"
	"log"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func InitDB() {
 
    connStr := "postgres://postgres:Postgres21@localhost:5432/bookdb" 

    // Connect to DB
    conn, err := pgx.Connect(context.Background(), connStr)
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    DB = conn
    log.Println("Connected to PostgreSQL!")
}

func CloseDB() {
    DB.Close(context.Background())
}