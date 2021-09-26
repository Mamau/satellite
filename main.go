package main

import (
	"satellite/internal/commands"
	"satellite/internal/updater"

	"github.com/gookit/color"

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
