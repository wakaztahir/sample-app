package api

import (
	"net/http"
)

// SigninHandler handles the requests for signins / logins
func (server *Server) signinHandler(w http.ResponseWriter,r *http.Request){
	server.configureResponse(w, r)
}
