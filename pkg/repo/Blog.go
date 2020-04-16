package repo

import (
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

// Blog : Post data structure
type Blog struct {
	gorm.Model         // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into the model
	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null;unique" json:"description"`
	Author      User   `json:"author"`
	Key         string   `"gorm: unique"json:"key"`
	Locale      string `gorm:"not null" json:"locale"`
	Posts       []Post `json:"posts"`
}

// Prepare : Prepare
func (b *Blog) Prepare() {
	b.ID = 0
	b.Name = html.EscapeString(strings.TrimSpace(b.Name))
	b.Description = html.EscapeString(strings.TrimSpace(b.Description))
	b.Author = User{}
	b.Locale = "en-GB"
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}
