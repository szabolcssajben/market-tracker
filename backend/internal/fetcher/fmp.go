package fetcher

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/szabolcssajben/market-tracker/internal/db"
)

func ParseFMPHistoricalJson(data []byte, symbol string) ([]db.MarketData, error) {
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

	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse reponse from FMP: %w", err)
	}

	var results []db.MarketData
	for _, h := range raw.Historical {
		ts, err := time.Parse("2006-01-02", h.Date)
		if err != nil {
			return nil, fmt.Errorf("invalid date format: %w", err)
		}

		results = append(results, db.MarketData{
			IndexName:  symbol,
			Region:     inferRegion(symbol),
			Currency:   inferCurrency(symbol),
			Timestamp:  ts,
			OpenPrice:  h.Open,
			ClosePrice: h.Close,
			High:       h.High,
			Low:        h.Low,
			Volume:     h.Volume,
		})
	}

	return results, nil
}

// Map the symbols to regions
func inferRegion(symbol string) string {
	switch symbol {
	case "^GSPC":
		return "US"
	case "^N225":
		return "JP"
	case "^FTSE":
		return "EU"
	default:
		return "UNKNOWN"
	}
}

// Map the symbols to currencies
func inferCurrency(symbol string) string {
	switch symbol {
	case "^GSPC":
		return "USD"
	case "^N225":
		return "JPY"
	case "^FTSE":
		return "GBP"
	default:
		return "UNKNOWN"
	}
}
