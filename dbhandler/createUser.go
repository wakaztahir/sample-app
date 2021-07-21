package dbhandler

import (
	"SampleApp/models"
	"context"
	"database/sql"
	"errors"
	"time"
)

func CreateUser(db *sql.DB, user *models.User) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	query := `INSERT INTO Users (name,username,email,password,email_verified) VALUES ($1,$2,$3,$4,$5) RETURNING id`

	rows, err := db.QueryContext(ctx, query, user.Name, user.Username, user.Email, user.Password, 0)

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
