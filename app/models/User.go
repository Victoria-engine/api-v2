package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/Victoria-engine/api-v2/app/auth"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

// User : User data structure
type User struct {
	gorm.Model        // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into the model
	FirstName  string `gorm:"size:100;not null;" json:"first_name"`
	LastName   string `gorm:"size:100;not null;" json:"last_name"`
	Email      string `gorm:"size:150;not null;unique" json:"email"`
	Password   string `gorm:"size:100;not null;" json:"password"`
	BlogID     int    `json:"blog_id"`
}

// BeforeSave : BeforeSave GORM hook
func (u *User) BeforeSave() error {
	hashedPassword, err := auth.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

// Prepare : Initialzes a user model
func (u *User) Prepare() {
	u.ID = 0
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.BlogID = -1
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate : User model validations
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.FirstName == "" {
			return errors.New("Required FirstName")
		}
		if u.LastName == "" {
			return errors.New("Required LastName")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.FirstName == "" {
			return errors.New("Required First Name")
		}
		if u.LastName == "" {
			return errors.New("Required Last Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

// SaveUser : Creates a user into the db
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// FindAllUsers : Finds all users in the table
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}

	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}

	return &users, err
}

// FindUserByID : Finds a user by ID
func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}

	return u, err
}

// UpdateUser : Updates a user
func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"email":      u.Email,
			"update_at":  time.Now(),
		},
	)

	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// DeleteUser : Deletes a user
func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
