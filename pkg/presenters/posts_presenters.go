package presenters

import "github.com/Victoria-engine/api-v2/pkg/repo"

// SavePostResponse : SavePost Response
type SavePostResponse struct {
	Title   string `gorm:"size:255;not null;" json:"title"`
	Content string `gorm:"size:255;not null;" json:"content"`
}

// SavePostPresenter : SavePost Presenter
func SavePostPresenter(p repo.Post) SavePostResponse {
	return SavePostResponse{p.Title, p.Content}
}
