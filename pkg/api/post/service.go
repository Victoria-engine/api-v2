package post

import (
	"github.com/Victoria-engine/api-v2/pkg/api/post/repo"
	ur "github.com/Victoria-engine/api-v2/pkg/api/user/repo"
	"github.com/jinzhu/gorm"
)

// Service : Post service with their dependencies
type Service interface {
	SavePost(p repo.PostModel, owner ur.UserModel) (repo.PostModel, error)
}

type Repo interface {
	SavePost(db *gorm.DB) (*repo.PostModel, error)
}

// Post
type Post struct {
	db   *gorm.DB
	repo Repo
	// roleBaseAccessControl ...etc
}

// Initialize : Initializes a new Service
func Initialize(db *gorm.DB, repo Repo) Post {
	return *New(db, repo)
}

func New(db *gorm.DB, repo Repo) *Post {
	return &Post{db, repo}
}
