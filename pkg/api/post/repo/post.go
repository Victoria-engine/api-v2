package repo

import (
	"errors"
	ur "github.com/Victoria-engine/api-v2/pkg/api/user/repo"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

// Post : Post data structure
type PostModel struct {
	gorm.Model              // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into the model
	Title      string       `gorm:"size:255;not null;" json:"title"`
	Content    string       `gorm:"size:255;not null;" json:"content"`
	Author     ur.UserModel `json:"author"`
	BlogID     uint         `gorm:"not null"`
}

// Prepare : Prepare
func (p *PostModel) Prepare() {
	p.ID = 0
	p.BlogID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = ur.UserModel{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// Validate : Validations
func (p *PostModel) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	return nil
}

// SavePost : Saves a post
func (p *PostModel) SavePost(db *gorm.DB) (*PostModel, error) {
	var err error
	err = db.Debug().Model(&PostModel{}).Create(&p).Error
	if err != nil {
		return &PostModel{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&ur.UserModel{}).Where("id = ?", p.Author.ID).Take(&p.Author).Error
		if err != nil {
			return &PostModel{}, errors.New("cannot attach post to user because user does not exist " + err.Error())
		}
	}

	return p, nil
}

// FindAllPosts : Finds all posts in the table
func (p *PostModel) FindAllPosts(db *gorm.DB) (*[]PostModel, error) {
	var err error
	var posts []PostModel

	err = db.Debug().Model(&PostModel{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]PostModel{}, err
	}

	if len(posts) > 0 {
		for i := range posts {
			err := db.Debug().Model(&ur.UserModel{}).Where("id = ?", posts[i].Author.ID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]PostModel{}, err
			}
		}
	}
	return &posts, nil
}

// FindPostByID : Find a post by ID
func (p *PostModel) FindPostByID(db *gorm.DB, pid uint64) (*PostModel, error) {
	var err error
	err = db.Debug().Model(&PostModel{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &PostModel{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&ur.UserModel{}).Where("id = ?", p.Author.ID).Take(&p.Author).Error
		if err != nil {
			return &PostModel{}, err
		}
	}
	return p, nil
}

// UpdatePost : Updates a post
func (p *PostModel) UpdatePost(db *gorm.DB) (*PostModel, error) {

	var err error

	err = db.Debug().Model(&PostModel{}).Where("id = ?", p.ID).Updates(
		PostModel{
			Title:   p.Title,
			Content: p.Content,
		}).Error

	if err != nil {
		return &PostModel{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&ur.UserModel{}).Where("id = ?", p.BlogID).Take(&p.Author).Error
		if err != nil {
			return &PostModel{}, err
		}
	}
	return p, nil
}

// DeletePost : Deletes a post
func (p *PostModel) DeletePost(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&PostModel{}).Where("id = ? and author_id = ?", pid, uid).Take(&PostModel{}).Delete(&PostModel{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
