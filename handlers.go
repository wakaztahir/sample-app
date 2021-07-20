package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// configureResponse configures the response according to app ServerConfig
func (app *App) configureResponse(w http.ResponseWriter, r *http.Request) {
	if app.config.allowCors {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		origin := r.Header.Get("Origin")
		for _, allowed := range app.config.corsAllowedFor {
			if allowed == origin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
		}
	}
}

type JsonError struct {
	Message string `json:"message"`
}

func (app *App) writeJson(w http.ResponseWriter, wrap string, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	err := encoder.Encode(wrapper)

	return err
}

// errorJson Returns error information in json format
func (app *App) errorJson(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	jsonError := JsonError{
		Message: message,
	}
	err := app.writeJson(w, "error", jsonError)
	if err != nil {
		log.Println("Error When Writing A Json Error", err)
	}
}
