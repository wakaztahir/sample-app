package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/**
Sets flags from the command to app configuration
*/
func (app *App) setFlags() {

}

type configuration struct {
	Modes []struct {
		Name       string `json:"name"`
		DbHost     string `json:"db-host,omitempty"`
		DbPort     int    `json:"db-port,omitempty"`
		DbName     string `json:"db-name,omitempty"`
		DbUser     string `json:"db-user,omitempty"`
		DbPassword string `json:"db-password,omitempty"`
	} `json:"modes"`
}

/**
Configures json from root config.json file if exists
*/
func (app *App) jsonConfigure() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Could not open config.json")
	}

	var config configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Could not read config.json")
	}

	for _, mode := range config.Modes {
		if mode.Name == string(app.config.mode) {

			//Configuring Database Parameters
			app.dbConfig.host = mode.DbHost
			app.dbConfig.port = mode.DbPort
			app.dbConfig.dbname = mode.DbName
			app.dbConfig.user = mode.DbUser
			app.dbConfig.password = mode.DbPassword
		}
	}

}
