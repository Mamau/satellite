package strategy

import (
	"github.com/mamau/satellite/config/docker"
	"github.com/mamau/satellite/libs"
)

type DaemonStrategy struct {
	docker *docker.Docker
}

func NewDaemonStrategy(config *docker.Docker) *DaemonStrategy {
	return &DaemonStrategy{
		docker: config,
	}
}

func (d *DaemonStrategy) ToCommand() []string {
	return libs.MergeSliceOfString([]string{
		d.docker.GetDockerCommand(),
		d.docker.GetDetached(),
		d.docker.GetCleanUp(),
		d.docker.GetEnvironmentVariables(),
		d.docker.GetPorts(),
		d.docker.GetDns(),
		d.docker.GetVolumes(),
		d.docker.GetImage(),
	})
}
