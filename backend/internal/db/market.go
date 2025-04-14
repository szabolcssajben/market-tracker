package db

import (
	"context"
	"fmt"
	"log"
	"os"
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

func InsertMarketDataBatch(conn *pgx.Conn, batch []MarketData, tb string, dryRun bool) error {
	table := tb
	if table == "" {
		table = os.Getenv("TEST_TABLE_NAME")
	}
	if table == "" {
		table = "market_data_test"
	}

	if dryRun {
		log.Printf("[DRY RUN] Would insert %d records into %s", len(batch), table)
		for _, d := range batch {
			log.Printf("[DRY] %s @ %s → %.2f %s", d.IndexName, d.Timestamp.Format("2006-01-02"), d.ClosePrice, d.Currency)
		}
		return nil
	}

	for _, d := range batch {
		err := InsertMarketData(conn, d, table)
		if err != nil {
			log.Printf("[ERROR] Failed to insert row: %v", err)
			return err
		}
	}
	return nil
}
