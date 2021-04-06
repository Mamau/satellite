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
	GetCacheVolume() string
	GetUserId() string
	GetEnvironmentVariables() string
	GetVersion() string
	SetVersion(string)
	GetHosts() string
	GetPorts() string
	GetVolumes() string
	GetDns() string
}

type Command struct {
	CmdName string
	Version string
	Image   string
	WorkDir string
	HomeDir string
	Args    []string
}

func (c *Command) GetClientCommand() string {
	cmd := c.CmdName

	if cmd == "" {
		return strings.Join(c.Args, " ")
	}

	mainCommand := append([]string{cmd}, c.Args...)
	return strings.Join(mainCommand, " ")
}

func (c *Command) GetImage() string {
	if c.Version != "" {
		return fmt.Sprintf("%s:%s", c.Image, c.Version)
	}
	return c.Image
}

func (c *Command) GetProjectVolume() string {
	return fmt.Sprintf("-v %s:%s", libs.GetPwd(), c.HomeDir)
}

func (c *Command) GetWorkDir() string {
	if c.WorkDir != "" {
		return fmt.Sprintf("--workdir=%s", c.WorkDir)
	}

	if c.HomeDir != "" {
		return fmt.Sprintf("--workdir=%s", c.HomeDir)
	}

	return ""
}
