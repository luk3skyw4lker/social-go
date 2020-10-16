package middlewares

import (
	"errors"
	"net/http"

	"github.com/luk3skyw4lker/social-go/api/auth"
	"github.com/luk3skyw4lker/social-go/api/responses"
)

// SetMiddlewareJSON is...
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		next(res, req)
	}
}

// SetMiddlewareAuth is...
func SetMiddlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		err := auth.TokenValid(req)

		if err != nil {
			responses.ERROR(res, http.StatusUnauthorized, errors.New("Unauthorized"))

			return
		}

		next(res, req)
	}
}
