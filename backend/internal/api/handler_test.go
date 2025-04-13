package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szabolcssajben/market-tracker/internal/db"
	"github.com/szabolcssajben/market-tracker/internal/testutils"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	HealthHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	// Assert the status code is 200
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetLatestMarket(t *testing.T) {
	testutils.LoadEnv(t)
	conn, err := db.ConnectDB()
	assert.NoError(t, err)
	defer conn.Close(context.Background())

	req := httptest.NewRequest("GET", "/api/markets/latest", nil)
	w := httptest.NewRecorder()

	// Call the yet-to-be-implemented handler
	GetLatestMarketsHandler(conn).ServerHTTP(w, req)
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
