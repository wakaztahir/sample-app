package api

import (
	"SampleApp/config"
	"SampleApp/db"
	mux2 "github.com/gorilla/mux"
)

type Server struct {
	router *mux2.Router
	config *config.ServerConfig
	handler     *db.Handler
}
