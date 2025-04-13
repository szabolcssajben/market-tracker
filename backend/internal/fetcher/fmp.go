package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/szabolcssajben/market-tracker/internal/db"
)

func FetchHistoricalData(symbol string, from, to time.Time) ([]db.MarketData, error) {
	apiKey := os.Getenv("FMP_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("FMP_API_KEY is not set")
	}

	fromStr := from.Format("2006-01-02")
	toStr := to.Format("2006-01-02")
	url := fmt.Sprintf(
		"https://financialmodelingprep.com/stable/historical-price-eod/full?symbol=%s&from=%s&to=%s&apikey=%s",
		symbol, fromStr, toStr, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from FMP: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return ParseFMPHistoricalJson(body, symbol)
}

func ParseFMPHistoricalJson(data []byte, symbol string) ([]db.MarketData, error) {
	var raw []struct {
		Symbol string  `json:"symbol"`
		Date   string  `json:"date"`
		Open   float64 `json:"open"`
		High   float64 `json:"high"`
		Low    float64 `json:"low"`
		Close  float64 `json:"close"`
		Volume int64   `json:"volume"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse response from FMP: %w", err)
	}

	var results []db.MarketData
	for _, h := range raw {
		ts, err := time.Parse("2006-01-02", h.Date)
		if err != nil {
			return nil, fmt.Errorf("invalid date format: %w", err)
		}

		results = append(results, db.MarketData{
			IndexName:  h.Symbol,
			Region:     inferRegion(h.Symbol),
			Currency:   inferCurrency(h.Symbol),
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
