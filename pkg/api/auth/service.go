package auth

import (
	"github.com/jinzhu/gorm"
)

// TokenGenerator represents token generator interface
type TokenGenerator interface {
	GenerateToken(uid uint) (string, error)
}

type Service interface {
	//Register(u *models.User) (string, error)
	Authenticate(email, password string) (string, error)
}

// Auth
type Auth struct {
	db *gorm.DB
	tokenGen TokenGenerator
}

// Initialize : Initializes a new Auth Service
func Initialize(db *gorm.DB, tokenGen TokenGenerator) Auth {
	return *New(db, tokenGen)
}

func New(db *gorm.DB, tokenGen TokenGenerator) *Auth {
	return &Auth{db, tokenGen}
}