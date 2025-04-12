package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/szabolcssajben/market-tracker/internal/api"
)

func main() {
	r := chi.NewRouter()
	r.Get("/health", api.HealthHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
