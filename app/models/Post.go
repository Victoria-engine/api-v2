package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Post : Post data structure
type Post struct {
	gorm.Model        // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into the model
	Title      string `gorm:"size:255;not null;" json:"title"`
	Content    string `gorm:"size:255;not null;" json:"content"`
	Author     User   `json:"author"`
	BlogID     uint   `gorm:"not null"`
}

// Prepare : Prepare
func (p *Post) Prepare() {
	p.ID = 0
	p.BlogID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// Validate : Validations
func (p *Post) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	return nil
}

// SavePost : Saves a post
func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.Author.ID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, errors.New("Cannot attach post to user becauser user does not exist " + err.Error())
		}
	}

	return p, nil
}

// FindAllPosts : Finds all posts in the table
func (p *Post) FindAllPosts(db *gorm.DB) (*[]Post, error) {
	var err error
	posts := []Post{}

	err = db.Debug().Model(&Post{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}

	if len(posts) > 0 {
		for i := range posts {
			err := db.Debug().Model(&User{}).Where("id = ?", posts[i].Author.ID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]Post{}, err
			}
		}
	}
	return &posts, nil
}

// FindPostByID : Find a post by ID
func (p *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.Author.ID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

// UpdatePost : Updates a post
func (p *Post) UpdatePost(db *gorm.DB) (*Post, error) {

	var err error

	err = db.Debug().Model(&Post{}).Where("id = ?", p.ID).Updates(
		Post{
			Title:   p.Title,
			Content: p.Content,
		}).Error

	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.BlogID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

// DeletePost : Deletes a post
func (p *Post) DeletePost(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Post{}).Where("id = ? and author_id = ?", pid, uid).Take(&Post{}).Delete(&Post{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
