package presenters

import (
	"time"

	"github.com/Victoria-engine/api-v2/app/models"
)

// GetUserInfoPresenter : GetUserInfo Presenter
func GetUserInfoPresenter(u *models.User) interface{} {
	data := struct {
		ID        uint       `json:"id"`
		FirstName string     `json:"first_name"`
		LastName  string     `json:"last_name"`
		Email     string     `json:"email"`
		BlogID    int        `json:"blog_id"`
		CreatedAt time.Time  `json:"created_at"`
		DeletedAt *time.Time `json:"deleted_at"`
		UpdatedAt time.Time  `json:"updated_at"`
	}{u.ID, u.FirstName, u.LastName, u.Email, u.BlogID, u.CreatedAt, u.DeletedAt, u.UpdatedAt}

	return data
}
