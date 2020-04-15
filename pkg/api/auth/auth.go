package auth

import (
	"github.com/Victoria-engine/api-v2/pkg/api/user/repo"
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

	var user repo.UserModel

	err = a.db.Debug().Model(repo.UserModel{}).Where("email = ?", email).Take(&user).Error
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
func (a Auth) Register(u repo.UserModel) (repo.UserModel, error) {

	freshUser, err := u.SaveUser(a.db)
	if err != nil {

		return *freshUser, err
	}

	return *freshUser, nil
}
