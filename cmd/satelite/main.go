package main

import (
	"fmt"
	"log"

	"github.com/mamau/satellite/internal/updater"

	"github.com/mamau/satellite/internal/commands"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Printf("Current version is %s\n", updater.Version)
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}

	commands.InitServiceCommand()
	commands.Execute()
}
