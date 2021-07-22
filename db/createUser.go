package db

import (
	"SampleApp/models"
	"context"
	"errors"
	"time"
)

func (handler *Handler) CreateUser(user *models.User) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	query := `INSERT INTO Users (name,username,email,password,email_verified) VALUES ($1,$2,$3,$4,$5) RETURNING id`

	rows, err := handler.Db.QueryContext(ctx, query, user.Name, user.Username, user.Email, user.Password, 0)

	if err != nil {
		return 0, err
	} else {
		if rows.Next() {
			var id int
			err := rows.Scan(&id)
			if err != nil {
				return 0, err
			}

			return id, nil
		} else {
			return 0, errors.New("user could not be inserted")
		}
	}
}
