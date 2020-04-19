package presenters

import (
	"time"

	"github.com/Victoria-engine/api-v2/pkg/repo"
)

// SaveBlogResponse : SaveBlog Response
type SaveBlogResponse struct {
	Name        string              `json:"title"`
	Description string              `json:"content"`
	CreatedAt   time.Time           `json: "created_at"`
	DeletedAt   *time.Time          `json: "deleted_at"`
	UpdatedAt   time.Time           `json: "updated_at"`
	Author      GetUserInfoResponse `json: "author"`
	AuthorID    uint                `json:"author_id"`
	APIKey      string              `json:"key"`
	Locale      string              `json:"locale"`
	Posts       []repo.Post         `json:"posts"`
}

// SaveBlogPresenter : SaveBlog Presenter
func SaveBlogPresenter(p repo.Blog) SaveBlogResponse {
	authorData := GetUserInfoResponse{
		BlogID:    p.Author.BlogID,
		CreatedAt: p.Author.CreatedAt,
		DeletedAt: p.Author.DeletedAt,
		Email:     p.Author.Email,
		FirstName: p.Author.FirstName,
		LastName:  p.Author.LastName,
		ID:        p.Author.ID,
		UpdatedAt: p.Author.UpdatedAt,
	}

	return SaveBlogResponse{
		p.Name,
		p.Description,
		p.CreatedAt,
		p.DeletedAt,
		p.UpdatedAt,
		authorData,
		p.AuthorID,
		p.APIKey,
		p.Locale,
		p.Posts,
	}
}
