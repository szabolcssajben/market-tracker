package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/szabolcssajben/market-tracker/internal/api"
	"github.com/szabolcssajben/market-tracker/internal/db"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	// Set up the router
	r := chi.NewRouter()
	r.Get("/health", api.HealthHandler)

	// Start the server
	log.Println("Starting server on :8080")
	go func() {
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Supabase connection test
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer conn.Close(context.Background())

	if err := db.PingDB(conn); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}
	log.Println("DB connection successful!")
}
