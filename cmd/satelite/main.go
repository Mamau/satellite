package main

import (
	"log"

	"github.com/mamau/satellite/internal/commands"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}
	commands.Execute()
}
