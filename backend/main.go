package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/szabolcssajben/market-tracker/internal/api"
	"github.com/szabolcssajben/market-tracker/internal/db"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	// Supabase connection
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer conn.Close(context.Background())

	if err := db.PingDB(conn); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}
	log.Println("DB connection successful!")

	// Set up the router
	r := chi.NewRouter()

	// Allow CORS for dev
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	// Routes
	r.Get("/health", api.HealthHandler)
	r.Get("/api/markets/latest", api.GetLatestMarketsHandler(conn))

	// Test historic data fetch
	// result, err := fetcher.FetchHistoricalData("^GSPC", time.Date(2025, 1, 20, 0, 0, 0, 0, time.UTC), time.Now().UTC())
	// if err != nil {
	// 	log.Printf("Fetch error: %v", err)
	// } else {
	// 	log.Printf("Fetched %d entries", len(result))
	// }

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server error: %v", err)
	}

	// Supabase connection test
	// conn, err := db.ConnectDB()
	// if err != nil {
	// 	log.Fatalf("DB connection error: %v", err)
	// }
	// defer conn.Close(context.Background())

	// if err := db.PingDB(conn); err != nil {
	// 	log.Fatalf("DB ping error: %v", err)
	// }
	// log.Println("DB connection successful!")

	// Test insert into db
	// err = db.InsertMarketData(conn, db.MarketData{
	// 	IndexName:  "S&P 500",
	// 	Region:     "US",
	// 	Currency:   "USD",
	// 	Timestamp:  time.Now().UTC(),
	// 	OpenPrice:  5000.00,
	// 	ClosePrice: 5050.00,
	// 	High:       5100.00,
	// 	Low:        4950.00,
	// 	Volume:     120000000,
	// })
	// if err != nil {
	// 	log.Fatalf("DB insert failed: %v", err)
	// }
	// log.Println("Mock market data inserted")
}
