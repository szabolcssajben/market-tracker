package fetcher

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Raw test data
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
