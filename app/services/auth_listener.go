package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Victoria-engine/api-v2/app/auth"
	"github.com/Victoria-engine/api-v2/app/models"
	"github.com/Victoria-engine/api-v2/app/presenters"
	"github.com/Victoria-engine/api-v2/app/responses"
	"github.com/Victoria-engine/api-v2/app/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

// Login : Login a user
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}

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

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = auth.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(uint32(user.ID))
}
