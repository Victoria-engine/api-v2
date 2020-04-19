package repo

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Blog : Post data structure
type Blog struct {
	gorm.Model         // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into the model
	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null;unique" json:"description"`
	Author      User   `json:"author" json:"author"`
	AuthorID    uint   `"gorm":"not null" json:"author_id"`
	APIKey      string `"gorm: unique"json:"api_key"`
	Locale      string `gorm:"not null" json:"locale"`
	Posts       []Post `gorm:"not null" json:"posts"`
}

// Prepare : Prepare
func (b *Blog) Prepare() {
	b.ID = 0
	b.Name = html.EscapeString(strings.TrimSpace(b.Name))
	b.Description = html.EscapeString(strings.TrimSpace(b.Description))
	b.Author = User{}
	b.Posts = []Post{}
	b.Locale = "en-GB"
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

// Validate : Validations
func (b *Blog) Validate() error {

	if b.Description == "" {
		return errors.New("required description")
	}
	if b.Name == "" {
		return errors.New("required name")
	}

	return nil
}

func (b *Blog) Save(db *gorm.DB) (Blog, error) {
	var err error
	err = db.Debug().Model(&Blog{}).Create(&b).Error
	if err != nil {
		return Blog{}, err
	}

	return *b, nil
}
