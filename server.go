package main

import (
	"crypto/tls"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// configureRoutesForApiV1 Configures Routes With Api Version 1
func (app *App) configureRoutesForApiV1(config *ServerConfig) {
	if config.mode == DevelopmentMode {
		app.router.HandleFunc("/v1/signup", app.SignupHandler).Methods(http.MethodPost)
		app.router.HandleFunc("/v1/signin", app.SigninHandler).Methods(http.MethodPost)
	} else {
		app.router.HandleFunc("/v1/signup", app.SignupHandler).Methods(http.MethodPost).Schemes("https")
		app.router.HandleFunc("/v1/signin", app.SigninHandler).Methods(http.MethodPost).Schemes("https")
	}
}

// RunServer Configure And Runs Server On App Configured Port
func (app *App) RunServer(config *ServerConfig) {

	app.router = mux2.NewRouter()

	if config.mode == DevelopmentMode {
		//todo
	} else {
		//todo Check Server Host mux.Host()
	}

	app.configureRoutesForApiV1(config)

	var tlsConfig *tls.Config

	if config.mode != DevelopmentMode {
		tlsConfig = &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}
	} else {
		tlsConfig = nil
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		WriteTimeout: 5 * time.Second,
		Handler:      app.router,
		TLSConfig:    tlsConfig,
	}

	if !config.useHttps {
		//Running HTTP Server
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("error on ListenAndServer", err)
		}
	} else {
		//Running HTTPS Server
		config.isRunning = true
		err := server.ListenAndServeTLS(config.certificate, config.key)
		if err != nil {
			config.isRunning = false
			log.Fatal("error on ListenAndServeTLS", err)
		}
	}
}
