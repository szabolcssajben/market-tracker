package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szabolcssajben/market-tracker/internal/db"
	"github.com/szabolcssajben/market-tracker/internal/testutils"
)

func TestGetLatestMarket(t *testing.T) {
	// Load environment variables from .env file
	testutils.LoadEnv(t)
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		t.Fatal("Missing DATABASE_URL environment variable")
	}

	conn, err := db.ConnectDB()
	assert.NoError(t, err)
	defer conn.Close(context.Background())

	req := httptest.NewRequest("GET", "/api/markets/latest", nil)
	w := httptest.NewRecorder()

	// Call the yet-to-be-implemented handler
	GetLatestMarketsHandler(conn).ServeHTTP(w, req)
	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var data []db.MarketData
	err = json.NewDecoder(resp.Body).Decode(&data)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	// Sanity check
	for _, row := range data {
		assert.NotEmpty(t, row.IndexName)
		assert.NotEmpty(t, row.Region)
		assert.NotEmpty(t, row.Currency)
		assert.NotZero(t, row.Timestamp)
		assert.NotZero(t, row.OpenPrice)
		assert.NotZero(t, row.ClosePrice)
		assert.NotZero(t, row.High)
		assert.NotZero(t, row.Low)
		assert.NotZero(t, row.Volume)
	}
}
