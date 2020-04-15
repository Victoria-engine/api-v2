package seed

import (
	"github.com/Victoria-engine/api-v2/pkg/api/post/repo"
	ur "github.com/Victoria-engine/api-v2/pkg/api/user/repo"
	"log"

	"github.com/jinzhu/gorm"
)

var users = []ur.UserModel{
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

var posts = []repo.PostModel{
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

	err := db.Debug().DropTableIfExists(repo.PostModel{}, &repo.PostModel{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&repo.PostModel{}, &repo.PostModel{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repo.PostModel{}).AddForeignKey("blog_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&repo.PostModel{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		posts[i].BlogID = users[i].BlogID
		posts[i].Author = users[i]

		err = db.Debug().Model(&repo.PostModel{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
