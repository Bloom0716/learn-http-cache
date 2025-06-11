package main

import (
	"log"

	"github.com/Bloom0716/learn-http-cache/internal/database"
	"github.com/Bloom0716/learn-http-cache/internal/router"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file: ", err)
	}
}

func main() {
	repository, err := database.NewRepository()
	if err != nil {
		log.Fatal("error initializing database: ", err)
	}

	router := router.NewRouter(repository)
	router.Run(":8080")
}
