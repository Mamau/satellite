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

func (d *DockerComposeStrategy) GetExecCommand() string {
	return string(DOCKER_COMPOSE)
}

func (d *DockerComposeStrategy) ToCommand() []string {
	return pkg.MergeSliceOfString([]string{
		d.ctx.GetConfig().GetPath(),
		d.ctx.GetConfig().GetProjectDirectory(),
		d.ctx.GetConfig().GetLogLevel(),
		d.ctx.GetConfig().GetProjectName(),
		d.ctx.GetConfig().GetVerbose(),
		d.ctx.GetConfig().GetDockerCommand(),
		d.ctx.GetConfig().GetDetached(),
		d.ctx.GetConfig().GetRemoveOrphans(),
	})
}

func (d *DockerComposeStrategy) GetContext() CommandContext {
	return d.ctx
}
