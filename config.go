package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/**
Sets flags from the command to app configuration
*/
func (appConfig *ServerConfig) setFlags() {

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
func (appConfig *ServerConfig) jsonConfigure() {
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
	appConfig.key = config.KeyFile
	appConfig.certificate = config.CertificateFile

	for _, mode := range config.Modes {
		if mode.Name == string(appConfig.mode) {
			//Configuring Server Parameters
			appConfig.useHttps = mode.UseHttps
			appConfig.port = mode.ServerPort

			//Configuring Database Parameters
			appConfig.db.host = mode.DbHost
			appConfig.db.port = mode.DbPort
			appConfig.db.dbname = mode.DbName
			appConfig.db.user = mode.DbUser
			appConfig.db.password = mode.DbPassword
		}
	}

}
