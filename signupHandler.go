package main

import (
	"encoding/json"
	"net/http"
)

type signupParams struct {
}

// SignupHandler handles the requests for signups / new accounts
func (app *App) SignupHandler(w http.ResponseWriter, r *http.Request) {
	app.configureResponse(w, r)

	var params *signupParams

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		app.errorJson(w, "Invalid signup parameters", http.StatusBadRequest)
		return
	}
}
