package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type MarketData struct {
	IndexName  string
	Region     string
	Currency   string
	Timestamp  time.Time
	OpenPrice  float64
	ClosePrice float64
	High       float64
	Low        float64
	Volume     int64
}

func InsertMarketData(conn *pgx.Conn, data MarketData) error {
	_, err := conn.Exec(
		context.Background(),
		`INSERT INTO market_data (index_name, region, currency, timestamp, open_price, close_price, high, low, volume)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		data.IndexName,
		data.Region,
		data.Currency,
		data.Timestamp,
		data.OpenPrice,
		data.ClosePrice,
		data.High,
		data.Low,
		data.Volume,
	)
	return err
}
