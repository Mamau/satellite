package strategy

import (
	"github.com/mamau/satellite/pkg"
)

type DaemonStrategy struct {
	ctx CommandContext
}

func NewDaemonStrategy(ctx CommandContext) *DaemonStrategy {
	return &DaemonStrategy{
		ctx: ctx,
	}
}

func (d *DaemonStrategy) ToCommand() []string {
	return pkg.MergeSliceOfString([]string{
		d.ctx.GetConfig().GetDockerCommand(),
		d.ctx.GetConfig().GetDetached(),
		d.ctx.GetConfig().GetCleanUp(),
		d.ctx.GetConfig().GetNetwork(),
		d.ctx.GetConfig().GetEnvironmentVariables(),
		d.ctx.GetConfig().GetPorts(),
		d.ctx.GetConfig().GetDns(),
		d.ctx.GetConfig().GetVolumes(),
		d.ctx.GetConfig().GetImage(),
	})
}

func (d *DaemonStrategy) GetContext() CommandContext {
	return d.ctx
}
