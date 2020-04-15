package repo

import (
	"errors"
	"github.com/Victoria-engine/api-v2/pkg/utl/crypt"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"html"
	"log"
	"strings"
	"time"
)

// UserModel : User Model data structure
type UserModel struct {
	gorm.Model        // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into the model
	FirstName  string `gorm:"size:100;not null;" json:"first_name"`
	LastName   string `gorm:"size:100;not null;" json:"last_name"`
	Email      string `gorm:"size:150;not null;unique" json:"email"`
	Password   string `gorm:"size:100;not null;" json:"password"`
	BlogID     uint   `gorm:"default:0;" json:"blog_id"`
}

// BeforeSave : BeforeSave GORM hook
func (u *UserModel) BeforeSave() error {
	hashedPassword, err := crypt.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

// Prepare : Initialzes a user model
func (u *UserModel) Prepare() {
	u.ID = 0
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.BlogID = 0
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate : User model validations
func (u *UserModel) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
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
func (u *UserModel) SaveUser(db *gorm.DB) (*UserModel, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &UserModel{}, err
	}
	return u, nil
}

// FindAllUsers : Finds all users in the table
func (u *UserModel) FindAllUsers(db *gorm.DB) (*[]UserModel, error) {
	var err error
	users := []UserModel{}

	err = db.Debug().Model(&UserModel{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]UserModel{}, err
	}

	return &users, err
}

// FindUserByID : Finds a user by ID
func (u *UserModel) FindUserByID(db *gorm.DB, uid uint) (*UserModel, error) {
	var err error
	err = db.Debug().Model(UserModel{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserModel{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &UserModel{}, errors.New("user Not Found")
	}

	return u, err
}

// UpdateUser : Updates a user
func (u *UserModel) UpdateUser(db *gorm.DB, uid uint32) (*UserModel, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&UserModel{}).Where("id = ?", uid).Take(&UserModel{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"email":      u.Email,
			"update_at":  time.Now(),
		},
	)

	if db.Error != nil {
		return &UserModel{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&UserModel{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserModel{}, err
	}
	return u, nil
}

// DeleteByID : Deletes a user by ID
func (u *UserModel) DeleteByID(db *gorm.DB, uid uint) (int64, error) {

	db = db.Debug().Model(&UserModel{}).Where("id = ?", uid).Take(&UserModel{}).Delete(&UserModel{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
