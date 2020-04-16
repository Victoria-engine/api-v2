package middlewares

import (
	"errors"
	"net/http"

	"github.com/Victoria-engine/api-v2/pkg/auth"
	"github.com/Victoria-engine/api-v2/pkg/utils/responses"
)

// SetMiddlewareJSON : Converts the requests to JSON format
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication : Checks for authenticity of the requests
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := auth.IsTokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		next(w, r)
	}
}
