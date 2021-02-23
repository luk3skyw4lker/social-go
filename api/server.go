package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/luk3skyw4lker/social-go/api/controllers"
	"github.com/luk3skyw4lker/social-go/api/seed"
)

var server = controllers.Server{}

// Run is...
func Run() {
	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting .env. Reason: %v", err)
	} else {
		fmt.Println("ENV values loaded")
	}

	server.Initalize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"))

	if os.Getenv("ENVIRONMENT") == "DEV" {
		seed.Load(server.DB)
	}

	serverRunning := server.Run(":8080")

	if serverRunning != nil {
		log.Fatalln("Server failed to start")
	}
}
