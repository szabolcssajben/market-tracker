package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func GetLatestDateForIndex(conn *pgx.Conn, symbol string, tb string) (time.Time, bool, error) {

	table := tb
	if table == "" {
		table = "market_data"
	}

	var latest time.Time

	err := conn.QueryRow(context.Background(), fmt.Sprintf(`
		SELECT timestamp
		FROM %s
		WHERE index_name = $1
		ORDER BY timestamp DESC
		LIMIT 1
	`, table), symbol).Scan(&latest)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return time.Time{}, false, nil // Not found
		}
		return time.Time{}, false, err
	}

	return latest, true, nil
}
