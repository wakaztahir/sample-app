package dbhandler

import (
	"context"
	"database/sql"
	"time"
)

func EmailExists(db *sql.DB, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	query := `SELECT id FROM Users WHERE email=$1`

	rows, err := db.QueryContext(ctx, query, email)
	if err != nil {
		return true, err
	}

	return rows.Next(), nil
}
