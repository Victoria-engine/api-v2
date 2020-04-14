package api

import (
	"flag"
	"fmt"
	"github.com/Victoria-engine/api-v2/pkg/api"
	"log"
	"os"

	"github.com/joho/godotenv"
)


// Run : Runs the REST API and setups flags
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


	err = api.Start(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		*port,
		)
	if err != nil {
		panic(err.Error())
	}
}
