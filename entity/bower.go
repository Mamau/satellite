package entity

import (
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
	return append(dockerConfig, b.ClientCommand()...)
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
		b.GetProjectVolume(),
		b.GetImage(),
	})
}

func (b *Bower) GetImage() string {
	return b.Image
}

func (b *Bower) ClientCommand() []string {
	var preCmd []string
	var postCmd []string

	if b.Config != nil {
		cmd := b.Config.GetPreCommands()
		if len(cmd) > 0 {
			cmd += ";"
		}

		preCmd = libs.MergeSliceOfString([]string{cmd})
		postCmd = libs.MergeSliceOfString([]string{b.Config.GetPostCommands()})
	}

	ccmd := b.GetClientCommand()
	if len(postCmd) > 0 {
		ccmd += ";"
	}

	clientCmd := libs.MergeSliceOfString([]string{ccmd})

	preCmd = append(preCmd, clientCmd...)
	preCmd = append(preCmd, postCmd...)

	return libs.DeleteEmpty(preCmd)
}
