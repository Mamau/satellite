package main

import (
	"github.com/gookit/color"
	"github.com/mamau/satellite/internal/commands"
	"github.com/mamau/satellite/internal/updater"

	"github.com/joho/godotenv"
)

func main() {
	color.Info.Printf("Current version is %s\n", updater.Version)
	if err := godotenv.Load(); err != nil {
		color.Warn.Println("no .env file")
	}

	commands.InitServiceCommand()
	commands.InitMacrosSubCommand()
	commands.Execute()
}
