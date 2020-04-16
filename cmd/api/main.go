package api

import (
	"flag"
	"fmt"
	"github.com/Victoria-engine/api-v2/pkg/utils/seed"
	"log"
	"os"

	"github.com/Victoria-engine/api-v2/pkg/listeners"
	"github.com/joho/godotenv"
)

var server = listeners.Server{}

// Run : Runs the REST API
func Run() {
	var (
		port = flag.String("port", os.Getenv("PORT"), "The server port")
	)

	flag.Parse()
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("Got the env values !")
	}

	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	// Seed in case you need dummy data
	seed.Load(server.DB)

	server.Run(fmt.Sprintf(":%s", *port))
}
