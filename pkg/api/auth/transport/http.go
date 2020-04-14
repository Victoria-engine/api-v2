package transport

import (
	"encoding/json"
	"github.com/Victoria-engine/api-v2/pkg/api/auth"
	"github.com/Victoria-engine/api-v2/pkg/models"
	"github.com/Victoria-engine/api-v2/pkg/utl/formaterror"
	"github.com/Victoria-engine/api-v2/pkg/utl/responses"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// HTTP represents authmiddleware http service
type HTTP struct {
	svc auth.Service
}

func NewHTTP(svc auth.Service, mux *mux.Router) {
	h := HTTP{svc}

	mux.HandleFunc("/api/auth/login", h.login)

}

func (h *HTTP) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Method not allowed
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	u := models.User{}

	err = json.Unmarshal(body, &u)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err :=  h.svc.Authenticate(u.Email, u.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnauthorized, formattedError)
		return
	}

	res := auth.LoginPresenter(token)

	responses.JSON(w, http.StatusOK, res)

}
