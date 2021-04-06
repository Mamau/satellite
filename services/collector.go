package services

import (
	"github.com/mamau/starter/entity"

	"github.com/mamau/starter/libs"
)

type Collector struct {
	entity        Collectorable
	commandConfig *entity.Command
}

func NewCollector(c Collectorable) *Collector {
	return &Collector{
		entity: c,
	}
}

func (c *Collector) DockerConfigCommand() []string {
	var userId,
		workDir,
		envVars,
		imgVersion,
		hosts,
		ports,
		dns,
		volumes string

	if c.entity.GetDockerConfig() != nil {
		userId = c.entity.GetDockerConfig().GetUserId()
		workDir = c.entity.GetDockerConfig().GetWorkDir()
		envVars = c.entity.GetDockerConfig().GetEnvironmentVariables()
		imgVersion = c.entity.GetDockerConfig().GetVersion()
		hosts = c.entity.GetDockerConfig().GetHosts()
		ports = c.entity.GetDockerConfig().GetPorts()
		volumes = c.entity.GetDockerConfig().GetVolumes()
		dns = c.entity.GetDockerConfig().GetDns()
	}

	if imgVersion != "" {
		c.entity.GetDockerConfig().Version = imgVersion
	}

	if workDir == "" {
		workDir = c.entity.GetCommandConfig().GetWorkDir()
	}

	return libs.MergeSliceOfString([]string{
		userId,
		envVars,
		hosts,
		ports,
		dns,
		workDir,
		volumes,
		c.entity.GetProjectVolume(),
		c.entity.GetImage(),
	})
}

func (c *Collector) ClientCommand() []string {
	var preCmd []string
	var postCmd []string

	if c.entity.GetDockerConfig() != nil {
		cmd := c.entity.GetDockerConfig().GetPreCommands()
		if len(cmd) > 0 {
			cmd += ";"
		}

		preCmd = libs.MergeSliceOfString([]string{cmd})
		postCmd = libs.MergeSliceOfString([]string{c.entity.GetDockerConfig().GetPostCommands()})
	}

	ccmd := c.entity.GetCommandConfig().GetClientCommand()
	if len(postCmd) > 0 {
		ccmd += ";"
	}

	clientCmd := libs.MergeSliceOfString([]string{ccmd})

	preCmd = append(preCmd, clientCmd...)
	preCmd = append(preCmd, postCmd...)

	return libs.DeleteEmpty(preCmd)
}

func (c *Collector) CollectCommand() []string {
	return append(c.DockerConfigCommand(), c.entity.GetClientSignature(c.ClientCommand())...)
}

func (c *Collector) GetBeginCommand() []string {
	var bc []string

	if c.entity.GetDockerConfig() != nil {
		bc = append(bc, c.entity.GetDockerConfig().GetDockerCommand())
		bc = append(bc, c.entity.GetDockerConfig().GetFlags())
		bc = append(bc, c.entity.GetDockerConfig().GetDetached())
		bc = append(bc, c.entity.GetDockerConfig().GetCleanUp())
	} else {
		bc = append(bc, "run", "-ti")
	}

	return libs.DeleteEmpty(bc)
}
