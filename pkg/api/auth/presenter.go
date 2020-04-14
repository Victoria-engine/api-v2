package auth

import "github.com/Victoria-engine/api-v2/pkg/models"

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
	return LoginResponse{ token}
}

// RegisterPresenter : Register Presenter
func RegisterPresenter(p *models.User) RegisterResponse {
	return RegisterResponse{p.FirstName, p.LastName, p.Email}
}
