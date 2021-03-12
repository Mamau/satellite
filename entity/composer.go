package entity

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/mamau/starter/libs"
)

type Composer struct {
	Command
}

func (c *Composer) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		c.workDirVolume(),
		c.projectVolume(),
		c.certsVolume(),
		c.cacheVolume(),
		{c.getImage()},
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

func (c *Composer) certsVolume() []string {
	if runtime.GOOS != "windows" {
		return []string{}
	}

	return []string{
		"-v",
		"/etc/ssl/certs:/etc/ssl/certs",
	}
}
