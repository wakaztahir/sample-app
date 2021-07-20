package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	//Creating App Variable
	app := &App{
		config: &ServerConfig{
			isRunning: false,
			mode:      DevelopmentMode,
			db: &DatabaseConfig{

			},
		},
	}

	//Setting Flags
	app.config.setFlags()

	//Setting JSON Configuration
	app.config.jsonConfigure()

	//Openning Database Connection
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=20 sslmode=verify-full sslcert=cert/server.crt sslkey=cert/server.key sslrootcert=cert/server.crt",
		app.config.db.host,
		app.config.db.port,
		app.config.db.dbname,
		app.config.db.user,
		app.config.db.password)
	var err error
	app.db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error Opening Database", err)
	}

	//Running Server
	app.RunServer()
}
