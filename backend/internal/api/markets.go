package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/szabolcssajben/market-tracker/internal/db"
)

func GetLatestMarketsHandler(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := conn.Query(context.Background(),
			`SELECT DISTINCT ON (index_name)
			index_name, region, currency, timestamp,
			open_price, close_price, high, low, volume
		FROM market_data
		ORDER BY index_name, timestamp DESC
		`)

		if err != nil {
			http.Error(w, "Failed to fetch market data", http.StatusInternalServerError)
			return
		}

		defer rows.Close()
		var results []db.MarketData
		for rows.Next() {
			var row db.MarketData
			err := rows.Scan(
				&row.IndexName,
				&row.Region,
				&row.Currency,
				&row.Timestamp,
				&row.OpenPrice,
				&row.ClosePrice,
				&row.High,
				&row.Low,
				&row.Volume,
			)
			if err != nil {
				http.Error(w, "Failed to parse market data", http.StatusInternalServerError)
				return
			}
			results = append(results, row)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}
