package transport

import (
	"encoding/json"
	"fmt"
	"github.com/Victoria-engine/api-v2/pkg/api/auth"
	"github.com/Victoria-engine/api-v2/pkg/models"
	"github.com/Victoria-engine/api-v2/pkg/utl/formaterror"
	json2 "github.com/Victoria-engine/api-v2/pkg/utl/middleware/json"
	"github.com/Victoria-engine/api-v2/pkg/utl/responses"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// HTTP represents authmiddleware http service
type HTTP struct {
	// The Auth service
	svc auth.Service
}

func NewHTTP(svc auth.Service, mux *mux.Router) {
	h := HTTP{svc}

	authRoutes := mux.PathPrefix("/api/auth").Subrouter()

	authRoutes.HandleFunc("/login",
		json2.SetMiddlewareJSON(h.login))

	authRoutes.HandleFunc("/register",
		json2.SetMiddlewareJSON(h.register))

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

	err = u.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := h.svc.Authenticate(u.Email, u.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnauthorized, formattedError)
		return
	}

	res := auth.LoginPresenter(token)

	responses.JSON(w, http.StatusOK, res)
}

func (h *HTTP) register(w http.ResponseWriter, r *http.Request) {
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

	u.Prepare()

	err = u.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Save the user
	createdUser, err := h.svc.Register(u)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("user email already exists"))
		return
	}

	res := auth.RegisterPresenter(createdUser)

	responses.JSON(w, http.StatusOK, res)

}
