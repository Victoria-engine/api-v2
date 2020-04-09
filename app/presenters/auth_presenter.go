package presenters

import "github.com/Victoria-engine/api-v2/app/models"

// LoginResponse : Login response data
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

// RegisterResponse : Register response data
type RegisterResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// LoginPresenter : Login Presenter
func LoginPresenter(token string) LoginResponse {
	return LoginResponse{token}
}

// RegisterPresenter : Register Presenter
func RegisterPresenter(u *models.User) RegisterResponse {
	return RegisterResponse{u.FirstName, u.LastName, u.Email}
}
