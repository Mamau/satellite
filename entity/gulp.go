package entity

import (
	"fmt"
	"strings"

	"github.com/mamau/starter/libs"
)

type Gulp struct {
	WorkDir string
	HomeDir string
	Args    []string
}

func (g *Gulp) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		g.workDirVolume(),
		g.projectVolume(),
		g.getImage(),
		{g.fullCommand()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}

func (g *Gulp) fullCommand() string {
	return g.getConfigCommand() + g.getMainCommand()
}

func (g *Gulp) getConfigCommand() string {
	configCommand := libs.GetConfig().GetGulp().ToCommand()
	if configCommand != "" {
		configCommand += "; "
	}

	return configCommand
}

func (g *Gulp) getMainCommand() string {
	return strings.Join(g.Args, " ")
}

func (g *Gulp) getPostCommands() string {
	return fmt.Sprintf("; chown -R $USER_ID:$USER_ID %s", g.WorkDir)
}

func (g *Gulp) getImage() []string {
	return []string{
		fmt.Sprintf("mamau/gulp"),
	}
}

func (g *Gulp) workDirVolume() []string {
	if g.WorkDir == "" {
		g.WorkDir = g.HomeDir
	}
	return []string{
		fmt.Sprintf("--workdir=%s", g.WorkDir),
	}
}

func (g *Gulp) projectVolume() []string {
	return []string{
		"-v",
		fmt.Sprintf("%s:%s", libs.GetPwd(), g.HomeDir),
	}
}
