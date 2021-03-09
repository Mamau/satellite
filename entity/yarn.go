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
		{"/bin/bash", "-c", y.command()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}
func (y *Yarn) command() string {
	configArgs := y.getConfigArgs()
	if configArgs != "" {
		configArgs += "; "
	}
	fullCommand := configArgs + y.getArgs()
	return fullCommand
}
func (y *Yarn) getArgs() string {
	return strings.Join(y.Args, " ")
}
func (y *Yarn) getConfigArgs() string {
	return libs.GetConfig().GetYarn().ToCommand()
}
func (y *Yarn) getImage() []string {
	return []string{
		fmt.Sprintf("node:%s", y.Version),
	}
}
func (y *Yarn) workDirVolume() []string {
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
