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
		{"/bin/bash", "-c", c.command()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}
func (c *Composer) command() string {
	configArgs := c.getConfigArgs()
	if configArgs != "" {
		configArgs += "; "
	}
	fullCommand := configArgs + c.getArgs() + c.getPostCommands()
	return fullCommand
}
func (c *Composer) getPostCommands() string {
	return fmt.Sprintf("; chown -R $USER_ID:$USER_ID %s", c.WorkDir)
}
func (c *Composer) getConfigArgs() string {
	commands := []string{
		libs.GetConfig().GetComposer().GetRepository().ToCommand(),
	}
	return strings.Join(commands, "; ")
}
func (c *Composer) getArgs() string {
	return strings.Join(append(c.Args, "--ignore-platform-reqs"), " ")
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
