package controllers

import (
	"net/http"

	"github.com/Victoria-engine/api/app/responses"
)

// Home : Home
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
