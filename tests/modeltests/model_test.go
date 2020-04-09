package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Victoria-engine/api-v2/app/models"
	"github.com/Victoria-engine/api-v2/app/services"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = services.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}
var blogInstance = models.Blog{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	Database()

	os.Exit(m.Run())
}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			os.Getenv("TestDbHost"),
			os.Getenv("TestDbPort"),
			os.Getenv("TestDbUser"),
			os.Getenv("TestDbName"),
			os.Getenv("TestDbPassword"),
		)

		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("An error happened trying to connect to the Testing Database:", err)
		} else {
			fmt.Printf("Connected to the %s database\n", TestDbDriver)
		}
	}
}

// Refreshes both because posts depend on posts
func refreshUsersAndPostsTable() error {
	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed table")
	return nil
}

func refreshPostsTable() error {

	err := server.DB.DropTableIfExists(&models.Post{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.Post{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUser() (models.User, error) {
	user := models.User{
		FirstName: "Tester",
		LastName:  "User",
		Email:     "pet@gmail.com",
		Password:  "password",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedOnePost() (models.Post, error) {
	err := refreshPostsTable()
	if err != nil {
		log.Fatalln(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalln(err)
	}

	post := models.Post{
		Title:    "This is the title sam",
		Content:  "This is the content sam",
		AuthorID: user.ID,
		Author:   user,
	}

	err = server.DB.Create(&models.Post{}).Create(&post).Error
	if err != nil {
		log.Fatalln(err)
	}

	return post, nil
}

func seedUsers() ([]models.User, error) {
	users := []models.User{
		{
			FirstName: "Tester 1",
			LastName:  "User 1",
			Email:     "tester1@email.com",
			Password:  "password",
		},
		{
			FirstName: "Tester 2",
			LastName:  "User 2",
			Email:     "tester2@email.com",
			Password:  "password",
		},
	}

	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func seedPosts() ([]models.Post, error) {

	refreshPostsTable()

	var err error

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

	users, err := seedUsers()
	if err != nil {
		log.Fatalln(err)
	}

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}

	return posts, nil
}
