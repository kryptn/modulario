package main

import (
	"log"
	"github.com/joho/godotenv"

	client "github.com/kryptn/modulario/clients/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	app := client.BuildHandler()
	app()
}
