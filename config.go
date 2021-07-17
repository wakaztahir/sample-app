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
	CertificateFile string `json:"certificate_file"`
	KeyFile         string `json:"key_file"`
	Modes           []struct {
		Name       string `json:"name"`
		ServerPort int    `json:"server_port,omitempty"`
		DbHost     string `json:"db_host,omitempty"`
		DbPort     int    `json:"db_port,omitempty"`
		DbName     string `json:"db_name,omitempty"`
		DbUser     string `json:"db_user,omitempty"`
		DbPassword string `json:"db_password,omitempty"`
		UseHttps   bool   `json:"use_https"`
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

	//Configuring Global Parameters
	app.config.key = config.KeyFile
	app.config.certificate = config.CertificateFile

	for _, mode := range config.Modes {
		if mode.Name == string(app.config.mode) {
			//Configuring Server Parameters
			app.config.useHttps = mode.UseHttps
			app.config.port = mode.ServerPort

			//Configuring Database Parameters
			app.dbConfig.host = mode.DbHost
			app.dbConfig.port = mode.DbPort
			app.dbConfig.dbname = mode.DbName
			app.dbConfig.user = mode.DbUser
			app.dbConfig.password = mode.DbPassword
		}
	}

}
