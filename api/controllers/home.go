package controllers

import (
	"net/http"

	"github.com/kingwerd/fullstack-go-app/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to the Demo API")
}
