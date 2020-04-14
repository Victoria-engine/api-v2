package presenters

import (
	"time"

	"github.com/Victoria-engine/api-v2/pkg/models"
)

// GetUserInfoResponse : GetUserInfo Response data
type GetUserInfoResponse struct {
	ID        uint       `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	BlogID    uint       `json:"blog_id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// GetUserInfoPresenter : GetUserInfo Presenter
func GetUserInfoPresenter(u *models.User) GetUserInfoResponse {
	return GetUserInfoResponse{
		u.ID,
		u.FirstName,
		u.LastName,
		u.Email,
		u.BlogID,
		u.CreatedAt, u.DeletedAt, u.UpdatedAt}
}
