package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model        // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into the model
	Title      string `gorm:"size:255;not null;unique" json:"title"`
	Content    string `gorm:"size:255;not null;" json:"content"`
	Author     User   `json:"author"`
	AuthorID   uint32 `gorm:"not null" json:"author_id"`
}

func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}
