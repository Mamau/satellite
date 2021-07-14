package strategy

import (
	"strings"

	"github.com/mamau/satellite/pkg"
)

type RunStrategy struct {
	ctx  CommandContext
	Args []string
}

func NewRunStrategy(ctx CommandContext, args []string) *RunStrategy {
	return &RunStrategy{
		ctx:  ctx,
		Args: args,
	}
}

func (r *RunStrategy) GetExecCommand() string {
	return string(DOCKER)
}

func (r *RunStrategy) ToCommand() []string {
	bc := pkg.MergeSliceOfString([]string{
		r.ctx.GetConfig().GetDockerCommand(),
		r.ctx.GetConfig().GetFlags(),
		r.ctx.GetConfig().GetCleanUp(),
		r.ctx.GetConfig().GetUserId(),
		r.ctx.GetConfig().GetEnvironmentVariables(),
		r.ctx.GetConfig().GetHosts(),
		r.ctx.GetConfig().GetPorts(),
		r.ctx.GetConfig().GetDns(),
		r.ctx.GetConfig().GetWorkDir(),
		r.ctx.GetConfig().GetVolumes(),
		r.ctx.GetConfig().GetContainerName(),
		r.ctx.GetConfig().GetImage(),
	})

	return append(bc, r.clientCommand()...)
}

func (r *RunStrategy) clientCommand() []string {
	execCommand := r.ctx.GetConfig().GetImageCommand()
	isBinBash := strings.Contains(execCommand, "/bin/bash")

	preCommand := r.ctx.GetConfig().GetPreCommands()
	if len(preCommand) > 0 {
		preCommand[len(preCommand)-1] += ";"
	}

	clientCommand := r.getArgs()
	postCommand := r.ctx.GetConfig().GetPostCommands()
	if len(postCommand) > 0 {
		clientCommand[len(clientCommand)-1] += ";"
	}
	listCmd := append(preCommand, clientCommand...)
	clientCmd := append(listCmd, postCommand...)

	cleanExecCmd := pkg.DeleteEmpty(pkg.MergeSliceOfString([]string{execCommand}))

	if isBinBash {
		return pkg.DeleteEmpty(append(cleanExecCmd, strings.Join(clientCmd, " ")))
	}

	return pkg.DeleteEmpty(append(cleanExecCmd, clientCmd...))
}

func (r *RunStrategy) getArgs() []string {
	if len(r.ctx.GetConfig().GetPreCommands()) > 0 || len(r.ctx.GetConfig().GetPostCommands()) > 0 {
		return append([]string{r.ctx.GetConfig().ImageCommand}, r.Args...)
	}

	return r.Args
}

func (r *RunStrategy) GetContext() CommandContext {
	return r.ctx
}
