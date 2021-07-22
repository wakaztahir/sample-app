package db

import (
	"context"
	"time"
)

func (handler *Handler) EmailExists(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	query := `SELECT id FROM Users WHERE email=$1`

	rows, err := handler.Db.QueryContext(ctx, query, email)
	if err != nil {
		return true, err
	}

	return rows.Next(), nil
}
