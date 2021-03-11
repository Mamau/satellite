package entity

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/mamau/starter/libs"
)

type Composer struct {
	Version string
	WorkDir string
	HomeDir string
	Args    []string
}

func (c *Composer) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		c.workDirVolume(),
		c.projectVolume(),
		c.certsVolume(),
		c.cacheVolume(),
		c.getImage(),
		{"/bin/bash", "-c", c.fullCommand()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}

func (c *Composer) fullCommand() string {
	configArgs := c.getConfigCommand()
	if configArgs != "" {
		configArgs += "; "
	}
	fullCommand := c.getConfigCommand() + c.getMainCommand() + c.getPostCommands()
	return fullCommand
}

func (c *Composer) getPostCommands() string {
	return fmt.Sprintf("; chown -R $USER_ID:$USER_ID %s", c.WorkDir)
}

func (c *Composer) getConfigCommand() string {
	configCommand := libs.GetConfig().GetComposer().GetRepository().ToCommand()
	if configCommand != "" {
		configCommand += "; "
	}

	return configCommand
}

func (c *Composer) getMainCommand() string {
	mainCommand := append([]string{"composer"}, c.Args...)
	return strings.Join(append(mainCommand, "--ignore-platform-reqs"), " ")
}

func (c *Composer) cacheVolume() []string {
	return []string{
		"-v",
		fmt.Sprintf("%s/cache/composer:/tmp", libs.GetPwd()),
	}
}

func (c *Composer) getImage() []string {
	return []string{
		fmt.Sprintf("composer:%s", c.Version),
	}
}

func (c *Composer) certsVolume() []string {
	if runtime.GOOS != "windows" {
		return []string{}
	}

	return []string{
		"-v",
		"/etc/ssl/certs:/etc/ssl/certs",
	}
}

func (c *Composer) workDirVolume() []string {
	if c.WorkDir == "" {
		c.WorkDir = c.HomeDir
	}
	return []string{
		fmt.Sprintf("--workdir=%s", c.WorkDir),
	}
}

func (c *Composer) projectVolume() []string {
	return []string{
		"-v",
		fmt.Sprintf("%s:%s", libs.GetPwd(), c.HomeDir),
	}
}
