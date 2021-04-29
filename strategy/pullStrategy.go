package strategy

import (
	"github.com/mamau/satellite/config/docker"
	"github.com/mamau/satellite/libs"
)

type PullStrategy struct {
	docker *docker.Docker
}

func NewPullStrategy(config *docker.Docker) *PullStrategy {
	return &PullStrategy{
		docker: config,
	}
}

func (p *PullStrategy) ToCommand() []string {
	return libs.MergeSliceOfString([]string{
		p.docker.GetDockerCommand(),
		p.docker.GetImage(),
	})
}
