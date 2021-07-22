package db

import (
	"SampleApp/config"
	"database/sql"
	"fmt"
	"log"
)

func OpenDb(config *config.DatabaseConfig, sslEnabled bool, sslCert string, sslKey string, sslRootCert string) *Handler {
	certifyString := ""
	if sslEnabled {
		certifyString = fmt.Sprintf("sslmode=verify-full sslcert=%s sslkey=%s sslrootcert=%s", sslCert, sslKey, sslRootCert)
	} else {
		certifyString = "sslmode=disable"
	}

	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=20 %s",
		config.Host,
		config.Port,
		config.DbName,
		config.User,
		config.Password,
		certifyString,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error opening database", err)
	}

	return &Handler{
		Db:     db,
		config: config,
	}
}
