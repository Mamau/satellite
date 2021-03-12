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
		{b.fullCommand()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}

func (b *Bower) fullCommand() string {
	return b.getConfigCommand() + b.getMainCommand()
}

func (b *Bower) getConfigCommand() string {
	configCommand := libs.GetConfig().GetBower().ToCommand()
	if configCommand != "" {
		configCommand += "; "
	}

	return configCommand
}

func (b *Bower) getMainCommand() string {
	return strings.Join(b.Args, " ")
}

func (b *Bower) getPostCommands() string {
	return fmt.Sprintf("; chown -R $USER_ID:$USER_ID %s", b.WorkDir)
}

func (b *Bower) getImage() []string {
	return []string{
		fmt.Sprintf("mamau/bower"),
	}
}

func (b *Bower) workDirVolume() []string {
	if b.WorkDir == "" {
		b.WorkDir = b.HomeDir
	}
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
