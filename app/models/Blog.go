package models

import "github.com/jinzhu/gorm"

// Blog : Post data structure
type Blog struct {
	gorm.Model         // Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into the model
	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null;unique" json:"description"`
	Author      User   `json:"author"`
	AuthorID    uint   `gorm:"not null" json:"author_id"`
	Key         uint   `gorm:"not null" json:"key"`
	Locale      string `gorm:"not null" json:"locale"`
	Posts       []Post `json:"posts"`
}
