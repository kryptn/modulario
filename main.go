package main

import (
	//client "github.com/kryptn/modulario/clients/http"
	"github.com/joho/godotenv"

	client "github.com/kryptn/modulario/clients/http"

	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	//app := client.Handle
	//app()

	client.Handle()()
}
