package strategy

import (
	"github.com/mamau/satellite/pkg"
)

type PullStrategy struct {
	ctx CommandContext
}

func NewPullStrategy(ctx CommandContext) *PullStrategy {
	return &PullStrategy{
		ctx: ctx,
	}
}

func (p *PullStrategy) ToCommand() []string {
	return pkg.MergeSliceOfString([]string{
		p.ctx.GetConfig().GetDockerCommand(),
		p.ctx.GetConfig().GetImage(),
	})
}

func (p *PullStrategy) GetContext() CommandContext {
	return p.ctx
}
