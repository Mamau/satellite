package entity

import (
	"fmt"
	"strings"

	"github.com/mamau/starter/libs"
)

type DockerInterface interface {
	GetPreCommands() string
	SetPreCommands([]string)
	GetPostCommands() string
	SetPostCommands([]string)
	GetWorkDir() string
	SetWorkDir(string)
	GetCacheDir() string
	GetUserId() []string
	GetEnvironmentVariables() []string
	GetVersion() string
	SetVersion(string)
	GetHosts() []string
	GetPorts() []string
	GetVolumes() []string
	GetDns() []string
}

type Command struct {
	CmdName      string
	Version      string
	Image        string
	WorkDir      string
	HomeDir      string
	Dns          []string
	Args         []string
	DockerConfig DockerInterface
}

func (c *Command) getCommand() string {
	cmd := c.CmdName

	if cmd == "" {
		return strings.Join(c.Args, " ")
	}

	mainCommand := append([]string{cmd}, c.Args...)
	return strings.Join(mainCommand, " ")
}

func (c *Command) getImage() string {
	if cv := c.DockerConfig.GetVersion(); cv != "" {
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
	if cc := c.DockerConfig.GetPreCommands(); cc != "" {
		return cc + "; "
	}

	return ""
}

func (c *Command) getPostCommands() string {
	if pc := c.DockerConfig.GetPostCommands(); pc != "" {
		return "; " + pc
	}

	return ""
}

func (c *Command) workDir() []string {
	return []string{
		fmt.Sprintf("--workdir=%s", c.getWorkDir()),
	}
}

func (c *Command) getWorkDir() string {
	c.WorkDir = c.DockerConfig.GetWorkDir()
	if c.WorkDir == "" {
		c.WorkDir = c.HomeDir
	}
	return c.WorkDir
}

func (c *Command) cacheDir() []string {
	if c.DockerConfig.GetCacheDir() == "" {
		return nil
	}

	return []string{
		"-v",
		fmt.Sprintf("%s:/tmp", c.DockerConfig.GetCacheDir()),
	}
}

func (c *Command) projectVolume() []string {
	volumes := c.DockerConfig.GetVolumes()
	currentDir := []string{
		"-v",
		fmt.Sprintf("%s:%s", libs.GetPwd(), c.HomeDir),
	}

	volumes = append(volumes, currentDir...)
	return volumes
}

func (c *Command) dockerCommandData() [][]string {
	return [][]string{
		c.DockerConfig.GetUserId(),
		c.DockerConfig.GetEnvironmentVariables(),
		c.DockerConfig.GetHosts(),
		c.DockerConfig.GetPorts(),
		c.DockerConfig.GetDns(),
		c.workDir(),
		c.cacheDir(),
		c.projectVolume(),
		{c.getImage()},
	}
}

func (c *Command) dockerDataToCommand() []string {
	var fullCommand []string
	for _, command := range c.dockerCommandData() {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}
