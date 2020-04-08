package controllers

import (
	"net/http"

	"github.com/Victoria-engine/api-v2/app/responses"
)

// Home : Home
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
