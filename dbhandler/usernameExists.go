package dbhandler

import (
	"SampleApp/validate"
	"context"
	"database/sql"
	"time"
)

func UsernameExists(db *sql.DB, username string) (bool, error) {
	validation := validate.Username(username)
	if validation.IsValid {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		query := `SELECT id FROM Users WHERE username=$1`

		rows, err := db.QueryContext(ctx, query, username)
		if err != nil {
			return true, err
		}

		return rows.Next(), nil

	} else {
		return true, nil
	}
}
