package testutils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/Victoria-engine/api-v2/app/auth"
	"github.com/Victoria-engine/api-v2/app/middlewares"
	"github.com/Victoria-engine/api-v2/app/models"
	"github.com/Victoria-engine/api-v2/app/services"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server : The server instance
var Server = services.Server{}
var UserInstance = models.User{}
var PostInstance = models.Post{}
var BlogInstance = models.Blog{}

// Database : Initis the test Database
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

		Server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("An error happened trying to connect to the Testing Database:", err)
		} else {
			fmt.Printf("Connected to the %s database\n", TestDbDriver)
		}
	}
}

// AuthenticatedRequest : Makes an jwt authenticated request for protected routes
func AuthenticatedRequest(h http.HandlerFunc, method, url string, urlParams map[string]string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}

	req = mux.SetURLVars(req, urlParams)

	// Create a new token for the test user
	token, err := auth.CreateToken(1)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(
		middlewares.SetMiddlewareJSON(
			middlewares.SetMiddlewareAuthentication(h),
		),
	)

	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

// Request : Makes a normal request
func Request(h http.HandlerFunc, method, url string, urlParams map[string]string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal(err)
	}

	req = mux.SetURLVars(req, urlParams)

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(middlewares.SetMiddlewareJSON(h))
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

// RefreshUsersAndPostsTable : Refreshes both because posts depend on posts
func RefreshUsersAndPostsTable() error {
	err := Server.DB.DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}

	err = Server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed table")
	return nil
}

// RefreshPostsTable : RefreshPostsTable
func RefreshPostsTable() error {

	err := Server.DB.DropTableIfExists(&models.Post{}).Error
	if err != nil {
		return err
	}
	err = Server.DB.AutoMigrate(&models.Post{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed tables")
	return nil
}

// SeedOneUser : SeedOneUser
func SeedOneUser() (models.User, error) {
	user := models.User{
		FirstName: "Tester",
		LastName:  "User",
		Email:     "logged@email.com",
		Password:  "password",
	}

	err := Server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

// SeedOnePost : SeedOnePost
func SeedOnePost() (models.Post, error) {
	err := RefreshPostsTable()
	if err != nil {
		log.Fatalln(err)
	}

	user, err := SeedOneUser()
	if err != nil {
		log.Fatalln(err)
	}

	post := models.Post{
		Title:   "This is the title sam",
		Content: "This is the content sam",
		Author:  user,
		BlogID:  user.BlogID,
	}

	err = Server.DB.Create(&models.Post{}).Create(&post).Error
	if err != nil {
		log.Fatalln(err)
	}

	return post, nil
}

// SeedUsers : SeedUsers
func SeedUsers() ([]models.User, error) {
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
		err := Server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

// SeedPosts : SeedPosts
func SeedPosts() ([]models.Post, error) {

	RefreshPostsTable()

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

	users, err := SeedUsers()
	if err != nil {
		log.Fatalln(err)
	}

	for i := range users {
		err = Server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		posts[i].BlogID = users[i].BlogID
		posts[i].Author = users[i]

		err = Server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}

	return posts, nil
}
