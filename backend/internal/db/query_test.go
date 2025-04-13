package db

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLatestDateForIndex(t *testing.T) {
	conn := connectTestDB(t)
	defer conn.Close(context.Background())

	index := "^TEST"
	now := time.Now().UTC().Truncate(time.Second)
	table := os.Getenv("TEST_TABLE_NAME")

	// Insert mock entries
	err := InsertMarketData(conn, MarketData{
		IndexName:  index,
		Region:     "TEST",
		Currency:   "USD",
		Timestamp:  now,
		OpenPrice:  1000,
		ClosePrice: 1100,
		High:       1110,
		Low:        990,
		Volume:     123456,
	}, table)
	assert.NoError(t, err)

	latest, found, err := GetLatestDateForIndex(conn, index, table)
	assert.NoError(t, err)
	assert.True(t, found)
	assert.WithinDuration(t, now, latest, time.Second)
}
