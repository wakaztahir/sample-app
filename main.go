package main

import (
	"SampleApp/api"
	"SampleApp/config"
	"SampleApp/db"
	"SampleApp/smtp"
	_ "github.com/lib/pq"
)

func main() {

	//Creating App Variable
	appConfig := &config.Config{
		Mode: config.DevelopmentMode,
		Server: &config.ServerConfig{
			IsRunning: false,
		},
		Db: &config.DatabaseConfig{

		},
		Smtp: &config.SMTPConfig{

		},
	}

	//Configuring By Command Line Flags
	appConfig.SetFlags()

	//Configuring By Json Config File
	appConfig.ConfigureByJson("config.json")

	//Opening Database Connection
	handler := db.OpenDb(appConfig.Db, appConfig.UseHTTPS, appConfig.Certificate, appConfig.Key, appConfig.Certificate)

	//Setting Up Database
	handler.SetupDb()

	go smtp.RunServer(appConfig.Smtp)

	//Running Server
	api.RunServer(appConfig.Server, handler, appConfig.Smtp, appConfig.UseHTTPS, appConfig.Certificate, appConfig.Key)
}
