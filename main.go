package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/makhmudovazeez/sms-router/router"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Route()
}
