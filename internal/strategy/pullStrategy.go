package strategy

import (
	"context"

	"github.com/mamau/satellite/pkg"

	docker2 "github.com/mamau/satellite/internal/config/docker"
)

type PullStrategy struct {
	ctx    context.Context
	docker *docker2.Docker
}

func NewPullStrategy(ctx context.Context, config *docker2.Docker) *PullStrategy {
	return &PullStrategy{
		ctx:    ctx,
		docker: config,
	}
}

func (p *PullStrategy) ToCommand() []string {
	return pkg.MergeSliceOfString([]string{
		p.docker.GetDockerCommand(),
		p.docker.GetImage(),
	})
}

func (p *PullStrategy) GetContext() context.Context {
	return p.ctx
}
