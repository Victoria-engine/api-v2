package listeners

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Victoria-engine/api-v2/pkg/auth"
	"github.com/Victoria-engine/api-v2/pkg/presenters"
	"github.com/Victoria-engine/api-v2/pkg/repo"
	"github.com/Victoria-engine/api-v2/pkg/utils/formaterror"
	"github.com/Victoria-engine/api-v2/pkg/utils/responses"
	"golang.org/x/crypto/bcrypt"
)

// Login : Login a user
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Method not allowed
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := repo.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()

	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	res := presenters.LoginPresenter(token)

	responses.JSON(w, http.StatusOK, res)
}

// SignIn : Sign in a user
func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := repo.User{}

	err = server.DB.Debug().Model(repo.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = auth.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(uint32(user.ID))
}

// Register : Register a new user
func (server *Server) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Method not allowed
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := repo.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()

	err = user.Validate("")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	freshUser, err := user.SaveUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, fmt.Errorf("user email already exists"))
		return
	}

	res := presenters.RegisterPresenter(freshUser)

	responses.JSON(w, http.StatusOK, res)
}
