package presenters

import (
	"time"

	"github.com/Victoria-engine/api-v2/pkg/repo"
)

// SaveBlogResponse : SaveBlog Response
type SaveBlogResponse struct {
	ID          uint                `json:"id"`
	Name        string              `json:"title"`
	Description string              `json:"content"`
	CreatedAt   time.Time           `json: "created_at"`
	DeletedAt   *time.Time          `json: "deleted_at"`
	UpdatedAt   time.Time           `json: "updated_at"`
	Author      GetUserInfoResponse `json: "author"`
	AuthorID    uint                `json:"author_id"`
	APIKey      string              `json:"key"`
	Posts       []repo.Post         `json:"posts"`
}

// GetBlogResponse :: GetBlog Response
type GetBlogResponse struct {
	Name        string      `json:"title"`
	Description string      `json:"content"`
	Posts       []repo.Post `json:"posts"`
}

// SaveBlogPresenter : SaveBlog Presenter
func SaveBlogPresenter(b repo.Blog) SaveBlogResponse {
	authorData := GetUserInfoResponse{
		BlogID:    b.Author.BlogID,
		CreatedAt: b.Author.CreatedAt,
		DeletedAt: b.Author.DeletedAt,
		Email:     b.Author.Email,
		FirstName: b.Author.FirstName,
		LastName:  b.Author.LastName,
		ID:        b.Author.ID,
		UpdatedAt: b.Author.UpdatedAt,
	}

	return SaveBlogResponse{
		b.ID,
		b.Name,
		b.Description,
		b.CreatedAt,
		b.DeletedAt,
		b.UpdatedAt,
		authorData,
		b.AuthorID,
		b.APIKey,
		b.Posts,
	}
}

// GetBlogPresenter : GetBlog Presenter
func GetBlogPresenter(b repo.Blog) GetBlogResponse {
	return GetBlogResponse{
		b.Name,
		b.Description,
		b.Posts,
	}
}
