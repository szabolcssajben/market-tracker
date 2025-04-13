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
	table := os.Getenv("TEST_TABLE_NAME")

	data := MarketData{
		IndexName:  "^TEST",
		Region:     "EU",
		Currency:   "EUR",
		Timestamp:  time.Now().UTC(),
		OpenPrice:  1500.00,
		ClosePrice: 1550.00,
		High:       1600.00,
		Low:        1450.00,
		Volume:     1000000,
	}

	err := InsertMarketData(conn, data, table)
	assert.NoError(t, err)
}

func TestInsertMarketDataBatch(t *testing.T) {
	conn := connectTestDB(t)
	defer conn.Close(context.Background())
	table := os.Getenv("TEST_TABLE_NAME")

	// Clean slate
	_, err := conn.Exec(context.Background(), "TRUNCATE "+table)
	assert.NoError(t, err)

	now := time.Now().UTC().Truncate(time.Second)

	mockData := []MarketData{
		{
			IndexName:  "^BATCH",
			Region:     "TEST",
			Currency:   "USD",
			Timestamp:  now,
			OpenPrice:  100,
			ClosePrice: 110,
			High:       115,
			Low:        95,
			Volume:     123456,
		},
		{
			IndexName:  "^BATCH",
			Region:     "TEST",
			Currency:   "USD",
			Timestamp:  now.Add(-24 * time.Hour),
			OpenPrice:  98,
			ClosePrice: 108,
			High:       112,
			Low:        94,
			Volume:     123000,
		},
	}

	err = InsertMarketDataBatch(conn, mockData, table, false)
	assert.NoError(t, err)

	var count int
	err = conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM "+table+" WHERE index_name = '^BATCH'").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 2, count)
}
