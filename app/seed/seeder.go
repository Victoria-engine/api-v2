package seed

import (
	"log"

	"github.com/Victoria-engine/api-v2/app/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		FirstName: "Tiago",
		LastName:  "Ferreira",
		BlogID:    1,
		Email:     "tiago@gmail.com",
		Password:  "password",
	},
	{
		FirstName: "John",
		LastName:  "Doe",
		BlogID:    2,
		Email:     "jdoe@gmail.com",
		Password:  "password",
	},
}

var posts = []models.Post{
	{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

// Load : Loads dummy data into the database
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		posts[i].BlogID = users[i].BlogID
		posts[i].Author = users[i]

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
