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
			isRunning:   false,
			mode:        DevelopmentMode,
		},
		dbConfig: &DatabaseConfig{

		},
	}

	//Setting Command Line Flags
	app.setFlags()

	//Setting JSON Configuration
	app.jsonConfigure()

	//Openning Database Connection
	connStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=20 sslmode=verify-full sslcert=cert/server.crt sslkey=cert/server.key sslrootcert=cert/server.crt",
		app.dbConfig.host,
		app.dbConfig.port,
		app.dbConfig.dbname,
		app.dbConfig.user,
		app.dbConfig.password)
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error Opening Database", err)
	}

	//Running Server
	app.RunServer()
}
