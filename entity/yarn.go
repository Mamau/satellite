package entity

import (
	"fmt"
	"strings"

	"github.com/mamau/starter/libs"
)

type Yarn struct {
	Version string
	WorkDir string
	HomeDir string
	Args    []string
}

func (y *Yarn) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		y.workDirVolume(),
		y.projectVolume(),
		y.getImage(),
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

func (y *Yarn) getConfigCommand() string {
	configCommand := libs.GetConfig().GetYarn().ToCommand()
	if configCommand != "" {
		configCommand += "; "
	}
	return configCommand
}

func (y *Yarn) getPostCommands() string {
	return fmt.Sprintf("; chown -R $USER_ID:$USER_ID %s", y.WorkDir)
}

func (y *Yarn) getImage() []string {
	return []string{
		fmt.Sprintf("node:%s", y.Version),
	}
}

func (y *Yarn) workDirVolume() []string {
	if y.WorkDir == "" {
		y.WorkDir = y.HomeDir
	}
	return []string{
		fmt.Sprintf("--workdir=%s", y.WorkDir),
	}
}

func (y *Yarn) projectVolume() []string {
	return []string{
		"-v",
		fmt.Sprintf("%s:%s", libs.GetPwd(), y.HomeDir),
	}
}
