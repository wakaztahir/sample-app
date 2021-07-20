package main

import "net/http"

// SigninHandler handles the requests for signins / logins
func (app *App) SigninHandler(w http.ResponseWriter,r *http.Request){
	app.configureResponse(w, r)
}
