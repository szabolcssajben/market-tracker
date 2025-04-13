package db

import (
	"context"
	"fmt"
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

func InsertMarketData(conn *pgx.Conn, data MarketData, tb string) error {
	table := tb
	if table == "" {
		table = "market_data"
	}

	_, err := conn.Exec(
		context.Background(),
		fmt.Sprintf(
			`INSERT INTO %s (index_name, region, currency, timestamp, open_price, close_price, high, low, volume)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, table),
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
