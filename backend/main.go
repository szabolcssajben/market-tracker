package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/szabolcssajben/market-tracker/internal/api"
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
