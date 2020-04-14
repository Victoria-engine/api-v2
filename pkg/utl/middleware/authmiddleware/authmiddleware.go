package authmiddleware

import (
	"errors"
	"github.com/Victoria-engine/api-v2/pkg/utl/jwtauth"
	"github.com/Victoria-engine/api-v2/pkg/utl/responses"
	"net/http"
)

// SetMiddlewareAuthentication : Checks for authenticity of the requests
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := jwtauth.Service{}.IsTokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}

		next(w, r)
	}
}
