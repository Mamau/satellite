package entity

import (
	"strings"
	"sync"

	"github.com/mamau/starter/config"
	"github.com/mamau/starter/config/composer"

	"github.com/mamau/starter/libs"
)

var cOnce sync.Once
var cInstance *Composer

type Composer struct {
	Config *composer.Composer
	*Command
}

func NewComposer(version string, args []string) *Composer {
	cOnce.Do(func() {
		cInstance = &Composer{
			Config: config.GetConfig().GetComposer(),
			Command: &Command{
				CmdName: "composer",
				Image:   "composer",
				HomeDir: "/home/www-data",
				Version: version,
				Args:    args,
			},
		}
	})

	return cInstance
}

func (c *Composer) CollectCommand() []string {
	dockerConfig := c.dockerConfigCommand()
	clientCmd := []string{"/bin/bash", "-c", strings.Join(c.ClientCommand(), "; ")}
	return append(dockerConfig, clientCmd...)
}

func (c *Composer) dockerConfigCommand() []string {
	var userId,
		workDir,
		cacheVolume,
		envVars,
		imgVersion,
		hosts,
		ports,
		dns,
		volumes string

	if c.Config != nil {
		userId = c.Config.GetUserId()
		workDir = c.Config.GetWorkDir()
		cacheVolume = c.Config.GetCacheVolume()
		envVars = c.Config.GetEnvironmentVariables()
		imgVersion = c.Config.GetVersion()
		hosts = c.Config.GetHosts()
		ports = c.Config.GetPorts()
		volumes = c.Config.GetVolumes()
		dns = c.Config.GetDns()
	}
	if imgVersion != "" {
		c.Version = imgVersion
	}

	return libs.MergeSliceOfString([]string{
		userId,
		envVars,
		hosts,
		ports,
		dns,
		workDir,
		cacheVolume,
		volumes,
		c.GetImage(),
	})
}

func (c *Composer) ClientCommand() []string {
	var preCmd string
	var postCmd string

	if c.Config != nil {
		preCmd = c.Config.GetPreCommands()
		postCmd = c.Config.GetPostCommands()
	}

	return libs.DeleteEmpty([]string{
		c.configToCommand(),
		preCmd,
		c.GetClientCommand(),
		postCmd,
	})
}

func (c *Composer) configToCommand() string {
	configCommands := libs.DeleteEmpty(c.Config.GetAll())
	if all := configCommands; all != nil {
		return strings.Join(configCommands, "; ")
	}

	return ""
}
