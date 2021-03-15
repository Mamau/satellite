package entity

import (
	"fmt"
	"strings"

	"github.com/mamau/starter/libs"
)

type ConfigCommandInterface interface {
	ToCommand() string
	GetDns() []string
}

type Command struct {
	Version string
	Image   string
	WorkDir string
	HomeDir string
	Dns     []string
	Args    []string
	Config  ConfigCommandInterface
}

func (c *Command) getImage() string {
	if c.Version != "" {
		return fmt.Sprintf("%s:%s", c.Image, c.Version)
	}
	return c.Image
}

func (c *Command) fullCommand() string {
	return c.getConfigCommand() + c.getMainCommand()
}

func (c *Command) getMainCommand() string {
	return strings.Join(c.Args, " ")
}

func (c *Command) getConfigCommand() string {
	configCommand := c.Config.ToCommand()
	if configCommand != "" {
		configCommand += "; "
	}

	return configCommand
}

func (c *Command) getPostCommands() string {
	return fmt.Sprintf("; chown -R $USER_ID:$USER_ID %s", c.WorkDir)
}

func (c *Command) workDirVolume() []string {
	if c.WorkDir == "" {
		c.WorkDir = c.HomeDir
	}
	return []string{
		fmt.Sprintf("--workdir=%s", c.WorkDir),
	}
}

func (c *Command) projectVolume() []string {
	return []string{
		"-v",
		fmt.Sprintf("%s:%s", libs.GetPwd(), c.HomeDir),
	}
}
