package fetcher

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/szabolcssajben/market-tracker/internal/db"
)

const sampleFMPResponse = `
{
  "symbol": "^GSPC",
  "historical": [
    {
      "date": "2025-04-10",
      "open": 5100.00,
      "high": 5150.00,
      "low": 5080.00,
      "close": 5125.00,
      "volume": 45000000
    },
    {
      "date": "2025-04-09",
      "open": 5050.00,
      "high": 5100.00,
      "low": 5020.00,
      "close": 5085.00,
      "volume": 47000000
    }
  ]
}`

func TestParseFMPHistoricData(t *testing.T) {
	var raw struct {
		Symbol     string `json:"symbol"`
		Historical []struct {
			Date   string  `json:"date"`
			Open   float64 `json:"open"`
			High   float64 `json:"high"`
			Low    float64 `json:"low"`
			Close  float64 `json:"close"`
			Volume int64   `json:"volume"`
		} `json:"historical"`
	}

	err := json.Unmarshal([]byte(sampleFMPResponse), &raw)
	assert.NoError(t, err)

	var result []db.MarketData
	for _, h := range raw.Historical {
		ts, err := time.Parse("2006-01-02", h.Date)
		assert.NoError(t, err)

		result = append(result, db.MarketData{
			IndexName:  raw.Symbol,
			Region:     "US",
			Currency:   "USD",
			Timestamp:  ts,
			OpenPrice:  h.Open,
			ClosePrice: h.Close,
			High:       h.High,
			Low:        h.Low,
			Volume:     h.Volume,
		})
	}

	assert.Len(t, result, 2)
	assert.Equal(t, "^GSPC", result[0].IndexName)
	assert.Equal(t, float64(5100.00), result[0].OpenPrice)
	assert.Equal(t, time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC), result[0].Timestamp)
}
