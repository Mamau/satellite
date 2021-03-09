package entity

import (
	"fmt"
	"strings"

	"github.com/mamau/starter/libs"
)

type Bower struct {
	WorkDir string
	HomeDir string
	Args    []string
}

func (b *Bower) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		b.workDirVolume(),
		b.projectVolume(),
		b.getImage(),
		{b.command()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}
func (b *Bower) command() string {
	return b.getArgs()
}
func (b *Bower) getArgs() string {
	return strings.Join(b.Args, " ")
}

func (b *Bower) getImage() []string {
	return []string{
		fmt.Sprintf("mamau/bower"),
	}
}
func (b *Bower) workDirVolume() []string {
	return []string{
		fmt.Sprintf("--workdir=%s", b.WorkDir),
	}
}

func (b *Bower) projectVolume() []string {
	return []string{
		"-v",
		fmt.Sprintf("%s:%s", libs.GetPwd(), b.HomeDir),
	}
}
