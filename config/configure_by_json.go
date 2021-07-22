package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type jsonConfiguration struct {
	CertificateFile string `json:"certificate_file"`
	KeyFile         string `json:"key_file"`
	RecaptchaSecret string `json:"recaptcha_secret"`
	SMTP            struct {
		Registered []struct {
			EmailID string `json:"email-id"`
		} `json:"registered"`
	} `json:"smtp"`
	Modes []struct {
		Name             string   `json:"name"`
		ServerPort       int      `json:"server_port,omitempty"`
		DbHost           string   `json:"db_host,omitempty"`
		DbPort           int      `json:"db_port,omitempty"`
		DbName           string   `json:"db_name,omitempty"`
		DbUser           string   `json:"db_user,omitempty"`
		DbPassword       string   `json:"db_password,omitempty"`
		UseHTTPS         bool     `json:"use_https"`
		AllowCors        bool     `json:"allow_cors"`
		CorsAllowedHosts []string `json:"cors_allowed_hosts"`
		VerifyRecaptcha  bool     `json:"verify_recaptcha"`
	} `json:"modes"`
}

// ConfigureByJson configures the server config object by the given config.json
func (config *Config) ConfigureByJson(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Could not open config.json",err)
	}

	var jsonConfig jsonConfiguration
	err = json.Unmarshal(file, &jsonConfig)
	if err != nil {
		log.Fatal("Could not read config.json")
	}

	//Configuring Global Parameters
	config.Key = jsonConfig.KeyFile
	config.Certificate = jsonConfig.CertificateFile
	config.Server.RecaptchaSecret = jsonConfig.RecaptchaSecret

	for _, mode := range jsonConfig.Modes {
		if mode.Name == string(config.Mode) {

			//Configuring Global Parameters
			config.UseHTTPS = mode.UseHTTPS

			//Configuring Server Parameters
			config.Server.Port = mode.ServerPort
			config.Server.AllowCORS = mode.AllowCors
			config.Server.CORSAllowedFor = mode.CorsAllowedHosts
			config.Server.VerifyRecaptcha = mode.VerifyRecaptcha

			//Configuring Database Parameters
			config.Db.Host = mode.DbHost
			config.Db.Port = mode.DbPort
			config.Db.DbName = mode.DbName
			config.Db.User = mode.DbUser
			config.Db.Password = mode.DbPassword
		}
	}
}
