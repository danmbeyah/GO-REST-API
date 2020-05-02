package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go-rest/api/controllers"
	"go-rest/api/seed"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables %v", err)
	} else {
		fmt.Println("Environment variables loaded successfully")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")
}