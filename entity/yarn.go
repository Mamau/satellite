package entity

import (
	"strings"
)

type Yarn struct {
	Command
}

func (y *Yarn) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		y.workDirVolume(),
		y.projectVolume(),
		{y.getImage()},
		{"/bin/bash", "-c", y.fullCommand()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}

func (y *Yarn) fullCommand() string {
	return y.getConfigCommand() + y.getMainCommand() + y.getPostCommands()
}

func (y *Yarn) getMainCommand() string {
	mainCommand := append([]string{"yarn"}, y.Args...)
	return strings.Join(mainCommand, " ")
}
