package db

import (
	"context"
	"time"
)

func (handler *Handler) UsernameExists(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	query := `SELECT id FROM Users WHERE username=$1`

	rows, err := handler.Db.QueryContext(ctx, query, username)
	if err != nil {
		return true, err
	}

	return rows.Next(), nil
}
