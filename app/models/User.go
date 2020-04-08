package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User : User data structure
type User struct {
	gorm.Model        // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into the model
	FirstName  string `gorm:"size:100;not null;" json:"first_name"`
	LastName   string `gorm:"size:100;not null;" json:"last_name"`
	Email      string `gorm:"size:150;not null;unique" json:"email"`
	Password   string `gorm:"size:100;not null;" json:"password"`
	BlogID     uint32 `json:"blog_id"`
}

// HashPassword : Hashes the a password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword : Verifies a password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
