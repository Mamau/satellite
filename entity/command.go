package entity

import (
	"fmt"
	"strings"

	"github.com/mamau/starter/libs"
)

type DockerInterface interface {
	GetCommand() string
	GetPreCommands() string
	GetPostCommands() string
	GetWorkDir() string
	GetUserId() []string
	GetEnvironmentVariables() []string
	GetVersion() string
	GetHosts() []string
	GetPorts() []string
	GetVolumes() []string
	GetDns() []string
}

type Command struct {
	Name    string
	Version string
	Image   string
	WorkDir string
	HomeDir string
	Dns     []string
	Args    []string
	Config  DockerInterface
}

func (c *Command) getCommand() string {
	cmd := c.Config.GetCommand()
	if cmd == "" {
		cmd = c.Name
	}

	if cmd == "" {
		return strings.Join(c.Args, " ")
	}

	mainCommand := append([]string{cmd}, c.Args...)
	return strings.Join(mainCommand, " ")
}

func (c *Command) getImage() string {
	if cv := c.Config.GetVersion(); cv != "" {
		c.Version = cv
	}

	if c.Version != "" {
		return fmt.Sprintf("%s:%s", c.Image, c.Version)
	}
	return c.Image
}

func (c *Command) fullCommand() string {
	return c.getPreCommands() + c.getCommand() + c.getPostCommands()
}

func (c *Command) getPreCommands() string {
	if cc := c.Config.GetPreCommands(); cc != "" {
		return cc + "; "
	}

	return ""
}

func (c *Command) getPostCommands() string {
	if pc := c.Config.GetPostCommands(); pc != "" {
		return "; " + pc
	}
	if c.Config.GetUserId() != nil {
		return ""
	}
	return fmt.Sprintf("; chown -R $USER_ID:$USER_ID %s", c.WorkDir)
}

func (c *Command) workDir() []string {
	c.WorkDir = c.Config.GetWorkDir()
	if c.WorkDir == "" {
		c.WorkDir = c.HomeDir
	}
	return []string{
		fmt.Sprintf("--workdir=%s", c.WorkDir),
	}
}

func (c *Command) projectVolume() []string {
	volumes := c.Config.GetVolumes()
	currentDir := []string{
		"-v",
		fmt.Sprintf("%s:%s", libs.GetPwd(), c.HomeDir),
	}

	volumes = append(volumes, currentDir...)
	return volumes
}
func (c *Command) configCommandData() [][]string {
	return [][]string{
		c.Config.GetUserId(),
		c.Config.GetEnvironmentVariables(),
		c.Config.GetHosts(),
		c.Config.GetPorts(),
		c.Config.GetDns(),
		c.workDir(),
		c.projectVolume(),
		{c.getImage()},
		{"/bin/bash", "-c", c.fullCommand()},
	}
}

func (c *Command) CollectCommand() []string {
	var fullCommand []string
	for _, command := range c.configCommandData() {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}
