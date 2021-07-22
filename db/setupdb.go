package db

import (
	"context"
	"time"
)

func (handler *Handler) SetupDb() error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	//Creating Users Table
	query := `CREATE TABLE IF NOT EXISTS Users (
				id serial NOT NULL PRIMARY KEY,
				name VARCHAR(80),
				username VARCHAR(80),
				email VARCHAR(320),
    			password VARCHAR(100),
				email_verified INT
			)`

	_, err := handler.Db.QueryContext(ctx, query)

	return err

}
