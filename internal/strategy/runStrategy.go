package strategy

import (
	"context"
	"fmt"
	"github.com/mamau/satellite/internal/config/docker"
	"github.com/mamau/satellite/pkg"
)

type RunStrategy struct {
	ctx    context.Context
	docker *docker.Docker
	Args   []string
}

func NewRunStrategy(ctx context.Context, config *docker.Docker, args []string) *RunStrategy {
	return &RunStrategy{
		ctx:    ctx,
		docker: config,
		Args:   args,
	}
}

func (r *RunStrategy) ToCommand() []string {
	bc := pkg.MergeSliceOfString([]string{
		r.docker.GetDockerCommand(),
		r.docker.GetFlags(),
		r.docker.GetCleanUp(),
		r.docker.GetUserId(),
		r.docker.GetEnvironmentVariables(),
		r.docker.GetHosts(),
		r.docker.GetPorts(),
		r.docker.GetDns(),
		r.docker.GetWorkDir(),
		r.docker.GetVolumes(),
		r.docker.GetContainerName(),
		r.docker.GetImage(),
	})

	return append(bc, pkg.DeleteEmpty(r.clientCommand())...)
}

func (r *RunStrategy) clientCommand() []string {
	execCommand := r.docker.GetImageCommand()

	preCommand := r.docker.GetPreCommands()
	if len(preCommand) > 0 {
		preCommand[len(preCommand)-1] += ";"
	}

	clientCommand := r.getArgs()
	postCommand := r.docker.GetPostCommands()
	if len(postCommand) > 0 {
		clientCommand[len(clientCommand)-1] += ";"
	}
	listCmd := append(preCommand, clientCommand...)
	clientCmd := pkg.DeleteEmpty(append(listCmd, postCommand...))

	cleanExecCmd := pkg.DeleteEmpty(pkg.MergeSliceOfString([]string{execCommand}))

	fmt.Println(clientCmd, len(clientCmd))
	return append(cleanExecCmd, clientCmd...)
}

func (r *RunStrategy) getArgs() []string {
	if len(r.docker.GetPreCommands()) > 0 || len(r.docker.GetPostCommands()) > 0 {
		return append([]string{r.docker.ImageCommand}, r.Args...)
	}

	return r.Args
}

func (r *RunStrategy) GetContext() context.Context {
	return r.ctx
}
