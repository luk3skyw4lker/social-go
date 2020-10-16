package controllers

import (
	"net/http"

	"github.com/luk3skyw4lker/social-go/api/responses"
)

// Home is..
func (server *Server) Home(res http.ResponseWriter, req *http.Request) {
	responses.JSON(res, http.StatusOK, "Welcome to our API!")
}
