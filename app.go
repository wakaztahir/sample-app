package main

import "net/http"

type App struct {
	config   *ServerConfig
	dbConfig *DatabaseConfig
	server   *http.Server
}

type RunMode string

const (
	DevelopmentMode RunMode = "development"
	ProductionMode  RunMode = "production"
)

type ServerConfig struct {
	isRunning   bool
	mode        RunMode
	certificate string
	key         string
	port        int
	useHttps    bool
}

type DatabaseConfig struct {
	host     string
	port     int
	dbname   string
	user     string
	password string
}
