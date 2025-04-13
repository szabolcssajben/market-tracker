package fetcher

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Raw test data
const sampleFMPResponse = `
[
  {
    "symbol": "^GSPC",
    "date": "2025-04-10",
    "open": 5100,
    "high": 5150,
    "low": 5080,
    "close": 5125,
    "volume": 123456
  },
  {
    "symbol": "^GSPC",
    "date": "2025-04-10",
    "open": 32321.21,
    "high": 34639.39,
    "low": 32320.66,
    "close": 34609,
    "volume": 201600000,
    "change": 2287.79,
    "changePercent": 7.08,
    "vwap": 33472.565
  },
]`

// Test parsing the FMP response
func TestParseFMPHistoricData(t *testing.T) {
	result, err := ParseFMPHistoricalJson([]byte(sampleFMPResponse), "^GSPC")
	assert.NoError(t, err)

	assert.Len(t, result, 2)
	assert.Equal(t, "^GSPC", result[0].IndexName)
	assert.Equal(t, float64(5100.00), result[0].OpenPrice)
	assert.Equal(t, "US", result[0].Region)
	assert.Equal(t, "USD", result[0].Currency)
	assert.Equal(t, time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC), result[0].Timestamp)
}
