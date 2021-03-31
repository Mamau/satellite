package main

import (
	"log"

	"github.com/mamau/starter/cmd"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}
	cmd.Execute()
}
