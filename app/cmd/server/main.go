package main

import (
	"log"

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
	router := router.NewRouter()
	router.Run(":8080")
}
