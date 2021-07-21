package main

import (
	"database/sql"
	mux2 "github.com/gorilla/mux"
)

type App struct {
	router *mux2.Router
	db     *sql.DB
	config *ServerConfig
}

type RunMode string

const (
	DevelopmentMode RunMode = "development"
	ProductionMode  RunMode = "production"
)

type ServerConfig struct {
	isRunning       bool
	mode            RunMode
	certificate     string
	key             string
	port            int
	useHttps        bool
	allowCors       bool
	corsAllowedFor  []string
	db              *DatabaseConfig
	verifyRecaptcha bool
	recaptchaSecret string
}

type DatabaseConfig struct {
	host     string
	port     int
	dbname   string
	user     string
	password string
}
