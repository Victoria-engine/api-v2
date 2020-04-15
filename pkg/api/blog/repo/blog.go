package repo

import (
	pr "github.com/Victoria-engine/api-v2/pkg/api/post/repo"
	ur "github.com/Victoria-engine/api-v2/pkg/api/user/repo"
	"github.com/jinzhu/gorm"
)

// Blog : Post data structure
type Blog struct {
	gorm.Model                 // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into the model
	Name        string         `gorm:"size:255;not null;unique" json:"name"`
	Description string         `gorm:"size:255;not null;unique" json:"description"`
	Author      ur.UserModel   `json:"author"`
	Key         uint           `gorm:"not null" json:"key"`
	Locale      string         `gorm:"not null" json:"locale"`
	Posts       []pr.PostModel `json:"posts"`
}
