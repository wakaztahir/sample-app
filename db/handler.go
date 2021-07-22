package db

import (
	"SampleApp/config"
	"database/sql"
)

type Handler struct {
	Db     *sql.DB
	config *config.DatabaseConfig
}
