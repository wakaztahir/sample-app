package main

type App struct {
	config *ServerConfig
	dbConfig *DatabaseConfig
}

type RunMode string

const (
	DevelopmentMode RunMode = "development"
	ProductionMode RunMode = "production"
)

type ServerConfig struct {
	isRunning bool
	mode RunMode
	certificate string
	key string
}

type DatabaseConfig struct {
	host string
	port int
	dbname string
	user string
	password string
}
