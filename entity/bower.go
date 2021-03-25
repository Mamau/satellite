package entity

import (
	"strings"
	"sync"

	"github.com/mamau/starter/libs"

	"github.com/mamau/starter/config"
)

var bOnce sync.Once
var bInstance *Bower

type Bower struct {
	Config *config.Bower
	*Command
}

func NewBower(args []string) *Bower {
	bOnce.Do(func() {
		bInstance = &Bower{
			Config: config.GetConfig().GetBower(),
			Command: &Command{
				CmdName: "bower",
				Image:   "mamau/bower",
				HomeDir: "/home/node",
				Args:    args,
			},
		}
	})

	return bInstance
}

func (b *Bower) CollectCommand() []string {
	dockerConfig := b.dockerConfigCommand()
	clientCmd := []string{"/bin/bash", "-c", strings.Join(b.ClientCommand(), "; ")}
	return append(dockerConfig, clientCmd...)
}

func (b *Bower) dockerConfigCommand() []string {
	var userId,
		workDir,
		cacheVolume,
		envVars,
		imgVersion,
		hosts,
		ports,
		dns,
		volumes string

	if b.Config != nil {
		userId = b.Config.GetUserId()
		workDir = b.Config.GetWorkDir()
		cacheVolume = b.Config.GetCacheVolume()
		envVars = b.Config.GetEnvironmentVariables()
		imgVersion = b.Config.GetVersion()
		hosts = b.Config.GetHosts()
		ports = b.Config.GetPorts()
		volumes = b.Config.GetVolumes()
		dns = b.Config.GetDns()
	}
	if imgVersion != "" {
		b.Version = imgVersion
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
		b.GetImage(),
	})
}

func (b *Bower) ClientCommand() []string {
	var preCmd string
	var postCmd string

	if b.Config != nil {
		preCmd = b.Config.GetPreCommands()
		postCmd = b.Config.GetPostCommands()
	}

	return libs.DeleteEmpty([]string{
		preCmd,
		b.GetClientCommand(),
		postCmd,
	})
}
