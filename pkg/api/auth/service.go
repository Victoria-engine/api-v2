package auth

import (
	"github.com/Victoria-engine/api-v2/pkg/api/user/repo"
	"github.com/jinzhu/gorm"
)

// TokenGenerator represents token generator interface
type TokenGenerator interface {
	GenerateToken(uid uint) (string, error)
}

// Service : Auth service with their dependencies
type Service interface {
	Register(u repo.UserModel) (repo.UserModel, error)
	Authenticate(email, password string) (string, error)
}

// Auth
type Auth struct {
	db       *gorm.DB
	tokenGen TokenGenerator
}

// Initialize : Initializes a new Auth Service
func Initialize(db *gorm.DB, tokenGen TokenGenerator) Auth {
	return *New(db, tokenGen)
}

func New(db *gorm.DB, tokenGen TokenGenerator) *Auth {
	return &Auth{db, tokenGen}
}
