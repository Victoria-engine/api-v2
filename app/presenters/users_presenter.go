package presenters

import (
	"github.com/Victoria-engine/api-v2/app/models"
)

// GetUserInfoPresenter : GetUserInfo Presenter
func GetUserInfoPresenter(u *models.User) interface{} {
	data := struct {
		ID        uint   `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		BlogID    int    `json:"blog_id"`
	}{u.ID, u.FirstName, u.LastName, u.Email, u.BlogID}

	return data
}
