package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (app *App) RunServer() {

	mux := http.NewServeMux()

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

	if app.config.mode == DevelopmentMode {
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
