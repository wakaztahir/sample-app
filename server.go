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
func (app *App) configureRoutesForApiV1(router *mux2.Router) {

	router.HandleFunc("/signup", app.SignupHandler).Methods("POST")
	router.HandleFunc("/signin", app.SigninHandler).Methods("POST")
}

// RunServer Configure And Runs Server On App Configured Port
func (app *App) RunServer() {

	mux := mux2.NewRouter()

	app.configureRoutesForApiV1(mux)

	var tlsConfig *tls.Config

	if app.config.mode != DevelopmentMode {
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

	app.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		WriteTimeout: 5 * time.Second,
		Handler:      mux,
		TLSConfig:    tlsConfig,
	}

	if !app.config.useHttps {
		//Running HTTP Server
		err := app.server.ListenAndServe()
		if err != nil {
			log.Fatal("error on ListenAndServer", err)
		}
	} else {
		//Running HTTPS Server
		err := app.server.ListenAndServeTLS(app.config.certificate, app.config.key)
		if err != nil {
			log.Fatal("error on ListenAndServeTLS", err)
		}
	}
}
