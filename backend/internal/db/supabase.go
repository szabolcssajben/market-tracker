package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() (*pgx.Conn, error) {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Supabase: %w", err)
	}
	return conn, nil
}

func PingDB(conn *pgx.Conn) error {
	var result int
	err := conn.QueryRow(context.Background(), "SELECT 1").Scan(&result)
	if err != nil {
		return fmt.Errorf("DB ping failed: %w", err)
	}
	if result != 1 {
		return fmt.Errorf("unexpected ping result: %d", result)
	}
	return nil
}
