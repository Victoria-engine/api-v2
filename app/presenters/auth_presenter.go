package presenters

// LoginResponse : Login response data
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

// LoginPresenter : Login Presenter
func LoginPresenter(token string) LoginResponse {
	return LoginResponse{token}
}
