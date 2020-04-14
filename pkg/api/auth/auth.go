package auth

import (
	"github.com/Victoria-engine/api-v2/pkg/models"
	"github.com/Victoria-engine/api-v2/pkg/utl/crypt"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate : Authenticate a user
func (a Auth) Authenticate(email, password string) (string, error) {
	token, err := a.SignIn(email, password)
	if err != nil {
		return "", err
	}

	return token, nil
}

// SignIn : Sign in a user
func (a Auth) SignIn(email, password string) (string, error) {
	var err error

	user := models.User{}

	err = a.db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = crypt.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return a.tokenGen.GenerateToken(user.ID)
}

// Register : Register a new user
func (a Auth) Register(u models.User) (models.User, error) {

	freshUser, err := u.SaveUser(a.db)
	if err != nil {

		return *freshUser, err
	}

	return *freshUser, nil
}
