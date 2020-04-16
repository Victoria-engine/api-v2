package listeners

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Victoria-engine/api-v2/pkg/repo"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

// Server : server structure
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize : Initializes the server structure
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

		server.DB, err = gorm.Open(Dbdriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database!\n", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to the %s database! \n", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&repo.User{}, &repo.Post{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

// Run : Runs and serves the server instance
func (server *Server) Run(addr string) {
	fmt.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
