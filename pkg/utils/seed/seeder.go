package seed

import (
	"log"

	"github.com/Victoria-engine/api-v2/pkg/repo"
	"github.com/jinzhu/gorm"
)

var users = []repo.User{
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

var blog = repo.Blog{
	Name:        "Mighty Dev Blog",
	Description: "Where the adventure starts!",
	APIKey:      "adioj1m3913mdoasdol1dj1ld1jdlçasjdçld",
}

var posts = []repo.Post{
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

	err := db.Debug().DropTableIfExists(&repo.Post{}, &repo.User{}, &repo.Blog{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&repo.User{}, &repo.Post{}, &repo.Blog{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*	err = db.Debug().Model(&repo.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/
	for u := range users {
		err = db.Debug().Model(&repo.User{}).Create(&users[u]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	// Somehow this is making a Insertion...
	//blog.Posts = posts
	blog.Author = users[0]

	err = db.Debug().Model(&repo.Blog{}).Create(&blog).Error
	if err != nil {
		log.Fatalf("cannot seed blog table: %v", err)
	}

	for p := range posts {
		posts[p].BlogID = blog.ID
		posts[p].Author = users[0]

		err = db.Debug().Model(&repo.Post{}).Create(&posts[p]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
