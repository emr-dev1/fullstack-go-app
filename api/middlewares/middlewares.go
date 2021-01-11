package middlewares

import (
	"errors"
	"net/http"

	"github.com/kingwerd/fullstack-go-app/api/auth"
	"github.com/kingwerd/fullstack-go-app/api/responses"
)

// SetMiddlewareJSON sets the content type of the response to json.
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication checks if the token from the request is valid
// or writes an error to the HTTP response writer.
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
