package strategy

import (
	"context"

	"github.com/mamau/satellite/pkg"

	"github.com/mamau/satellite/internal/config/docker"
)

type DaemonStrategy struct {
	ctx    context.Context
	docker *docker.Docker
}

func NewDaemonStrategy(ctx context.Context, config *docker.Docker) *DaemonStrategy {
	return &DaemonStrategy{
		ctx:    ctx,
		docker: config,
	}
}

func (d *DaemonStrategy) ToCommand() []string {
	return pkg.MergeSliceOfString([]string{
		d.docker.GetDockerCommand(),
		d.docker.GetDetached(),
		d.docker.GetCleanUp(),
		d.docker.GetNetwork(),
		d.docker.GetEnvironmentVariables(),
		d.docker.GetPorts(),
		d.docker.GetDns(),
		d.docker.GetVolumes(),
		d.docker.GetImage(),
	})
}

func (d *DaemonStrategy) GetContext() context.Context {
	return d.ctx
}