package db

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/szabolcssajben/market-tracker/internal/testutils"
)

func connectTestDB(t *testing.T) *pgx.Conn {
	// Load environment variables from .env file
	testutils.LoadEnv(t)
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		t.Fatal("Missing DATABASE_URL environment variable")
	}

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		t.Fatal("Failed to connect to test DB:", err)
	}

	return conn
}

func TestInsertMarketData(t *testing.T) {
	conn := connectTestDB(t)
	defer conn.Close(context.Background())

	data := MarketData{
		IndexName:  "DAX",
		Region:     "EU",
		Currency:   "EUR",
		Timestamp:  time.Now().UTC(),
		OpenPrice:  1500.00,
		ClosePrice: 1550.00,
		High:       1600.00,
		Low:        1450.00,
		Volume:     1000000,
	}

	err := InsertMarketData(conn, data)
	assert.NoError(t, err)
}
