package strategy

import (
	"github.com/mamau/satellite/pkg"
)

type DockerComposeStrategy struct {
	ctx CommandContext
}

func NewDockerComposeStrategy(ctx CommandContext) *DockerComposeStrategy {
	return &DockerComposeStrategy{
		ctx: ctx,
	}
}

func (d *DockerComposeStrategy) ToCommand() []string {
	return pkg.MergeSliceOfString([]string{
		d.ctx.GetConfig().GetDockerCommand(),
		d.ctx.GetConfig().GetImage(),
	})
}

func (d *DockerComposeStrategy) GetContext() CommandContext {
	return d.ctx
}
