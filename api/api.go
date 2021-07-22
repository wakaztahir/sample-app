package api

import (
	"SampleApp/config"
	"SampleApp/db"
	"crypto/tls"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// configureRoutesForApiV1 Configures Routes With Api Version 1
func (server *Server) configureRoutesForApiV1(useHTTPS bool) {
	signupRoute := server.router.HandleFunc("/v1/signup", server.signupHandler).Methods(http.MethodPost)
	signInRoute := server.router.HandleFunc("/v1/signin", server.signinHandler).Methods(http.MethodPost)

	if useHTTPS {
		signupRoute.Schemes("https")
		signInRoute.Schemes("https")
	}
}

// RunServer Configure And Runs Server On App Configured Port
func RunServer(config *config.ServerConfig, handler *db.Handler, useHTTPS bool, certFile string, keyFile string) {

	router := mux2.NewRouter()

	server := &Server{
		router:  router,
		config:  config,
		handler: handler,
	}

	router.Use(mux2.CORSMethodMiddleware(router))

	server.configureRoutesForApiV1(useHTTPS)

	var tlsConfig *tls.Config

	if useHTTPS {
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

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		WriteTimeout: 5 * time.Second,
		Handler:      router,
		TLSConfig:    tlsConfig,
	}

	if !useHTTPS {
		//Running HTTP Server
		err := httpServer.ListenAndServe()
		if err != nil {
			log.Fatal("error on ListenAndServer", err)
		}
	} else {
		//Running HTTPS Server
		err := httpServer.ListenAndServeTLS(certFile, keyFile)
		if err != nil {
			log.Fatal("error on ListenAndServeTLS", err)
		}
	}
}
