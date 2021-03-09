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
		{g.command()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}
func (g *Gulp) command() string {
	return g.getArgs()
}
func (g *Gulp) getArgs() string {
	return strings.Join(g.Args, " ")
}

func (g *Gulp) getImage() []string {
	return []string{
		fmt.Sprintf("mamau/gulp"),
	}
}
func (g *Gulp) workDirVolume() []string {
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
