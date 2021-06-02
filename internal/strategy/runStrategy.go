package strategy

import (
	"context"
	"fmt"
	"strings"

	"github.com/mamau/satellite/pkg"

	docker2 "github.com/mamau/satellite/internal/config/docker"
)

type RunStrategy struct {
	ctx    context.Context
	docker *docker2.Docker
	Args   []string
}

func NewRunStrategy(ctx context.Context, config *docker2.Docker, args []string) *RunStrategy {
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
		preCommand += ";"
	}

	clientCommand := r.getArgs()
	postCommand := r.docker.GetPostCommands()
	if len(postCommand) > 0 {
		clientCommand += ";"
	}

	listCmd := []string{
		preCommand,
		clientCommand,
		postCommand,
	}
	clientCmd := fmt.Sprintf("%s", strings.Join(pkg.DeleteEmpty(listCmd), " "))
	cleanExecCmd := pkg.DeleteEmpty(pkg.MergeSliceOfString([]string{execCommand}))

	return append(cleanExecCmd, clientCmd)
}

func (r *RunStrategy) getArgs() string {
	if r.docker.ImageCommand == "" {
		return ""
	}

	if len(r.docker.GetPreCommands()) > 0 || len(r.docker.GetPostCommands()) > 0 {
		cmd := append([]string{r.docker.ImageCommand}, r.Args...)
		return strings.Join(cmd, " ")
	}

	return strings.Join(r.Args, " ")
}

func (r *RunStrategy) GetContext() context.Context {
	return r.ctx
}
