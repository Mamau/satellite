package entity

import (
	"fmt"
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
	clientCmd := []string{"/bin/bash", "-c", strings.Join(c.ClientCommand(), " ")}
	return append(dockerConfig, clientCmd...)
}

//TODO: Обратить внимание на workdir, чтобы можно было без конфига использовать
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
		c.GetProjectVolume(),
		c.GetImage(),
	})
}

func (c *Composer) ClientCommand() []string {
	mainCmd := c.configToCommand()
	var preCmd []string
	var postCmd []string

	if c.Config != nil {
		cmd := c.Config.GetPreCommands()
		if len(cmd) > 0 {
			cmd += ";"
		}

		preCmd = libs.MergeSliceOfString([]string{cmd})
		postCmd = libs.MergeSliceOfString([]string{c.Config.GetPostCommands()})
	}

	ccmd := c.GetClientCommand()
	if len(postCmd) > 0 {
		ccmd += ";"
	}

	clientCmd := libs.MergeSliceOfString([]string{ccmd})

	mainCmd = append(mainCmd, preCmd...)
	mainCmd = append(mainCmd, clientCmd...)
	mainCmd = append(mainCmd, postCmd...)

	return libs.DeleteEmpty(mainCmd)
}

func (c *Composer) configToCommand() []string {
	configCommands := libs.DeleteEmpty(c.Config.GetAll())
	if len(configCommands) == 0 {
		return []string{}
	}

	for i, v := range configCommands {
		configCommands[i] = v + ";"
	}

	return libs.MergeSliceOfString(configCommands)
}

func (c *Composer) GetImage() string {
	if v := c.Config.GetVersion(); v != "" {
		return fmt.Sprintf("%s:%s", c.Image, v)
	}
	return c.Command.GetImage()
}
