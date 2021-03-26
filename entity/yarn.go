package entity

import (
	"strings"
	"sync"

	"github.com/mamau/starter/config"
	"github.com/mamau/starter/libs"
)

var once sync.Once
var instance *Yarn

type Yarn struct {
	Config *config.Yarn
	*Command
}

func NewYarn(version string, args []string) *Yarn {
	once.Do(func() {
		instance = &Yarn{
			Config: config.GetConfig().GetYarn(),
			Command: &Command{
				CmdName: "yarn",
				Image:   "node",
				HomeDir: "/home/node",
				Version: version,
				Args:    args,
			},
		}
	})

	return instance
}

func (y *Yarn) CollectCommand() []string {
	dockerConfig := y.dockerConfigCommand()
	clientCmd := []string{"/bin/bash", "-c", strings.Join(y.ClientCommand(), " ")}
	return append(dockerConfig, clientCmd...)
}

func (y *Yarn) dockerConfigCommand() []string {
	var userId,
		workDir,
		cacheVolume,
		envVars,
		imgVersion,
		hosts,
		ports,
		dns,
		volumes string

	if y.Config != nil {
		userId = y.Config.GetUserId()
		workDir = y.Config.GetWorkDir()
		cacheVolume = y.Config.GetCacheVolume()
		envVars = y.Config.GetEnvironmentVariables()
		imgVersion = y.Config.GetVersion()
		hosts = y.Config.GetHosts()
		ports = y.Config.GetPorts()
		volumes = y.Config.GetVolumes()
		dns = y.Config.GetDns()
	}
	if imgVersion != "" {
		y.Version = imgVersion
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
		y.GetProjectVolume(),
		y.GetImage(),
	})
}

func (y *Yarn) ClientCommand() []string {
	var preCmd []string
	var postCmd []string

	if y.Config != nil {
		cmd := y.Config.GetPreCommands()
		if len(cmd) > 0 {
			cmd += ";"
		}

		preCmd = libs.MergeSliceOfString([]string{cmd})
		postCmd = libs.MergeSliceOfString([]string{y.Config.GetPostCommands()})
	}

	ccmd := y.GetClientCommand()
	if len(postCmd) > 0 {
		ccmd += ";"
	}

	clientCmd := libs.MergeSliceOfString([]string{ccmd})

	preCmd = append(preCmd, clientCmd...)
	preCmd = append(preCmd, postCmd...)

	return libs.DeleteEmpty(preCmd)
}
