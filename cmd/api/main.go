package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Victoria-engine/api-v2/app/services"
	"github.com/joho/godotenv"
)

var server = services.Server{}

// Run : Runs the REST API
func Run() {

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
	//seed.Load(server.DB)

	server.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
