package server

import (
	"fmt"
	"github.com/Victoria-engine/api-v2/pkg/models"
	"github.com/Victoria-engine/api-v2/pkg/utl/responses"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

// Server : server structure
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
	Logger *log.Logger
}

//TODO: Add config struct to params
func New(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *Server {
	var err error
	server := &Server{
		Logger: log.New(os.Stdout, "http: ", log.LstdFlags),
		Router: mux.NewRouter(),
	}

	// Init database
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			DbHost, DbPort, DbUser, DbName, DbPassword)

		server.DB, err = gorm.Open(Dbdriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database!\n", Dbdriver)
			log.Fatal("Error:", err)
		} else {
			fmt.Printf("Connected to the %s database! \n", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	server.Router.HandleFunc("/health", healthCheck)

	return server
}

// TODO: Is this needed ?
func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.Router.ServeHTTP(w, r)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "All good!")
}

// TODO: Add config struct to params
// Run : Runs and serves the server instance
func (server *Server) Run(addr string) {

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt)

	serverAddr := ":" + addr

	h := &http.Server{
		Addr: serverAddr,
		//TODO: Add these timeout values to config values
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      server,
	}

	go func() {
		server.Logger.Printf("Listening on http://0.0.0.0%s\n", serverAddr)

		if err := h.ListenAndServe(); err != nil {
			server.Logger.Fatal(err)
		}
	}()

	<-stop

	server.Logger.Println("\nGracefully shutting down the server...")
}
