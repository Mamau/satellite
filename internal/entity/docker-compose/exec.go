package docker_compose

import (
	"fmt"
	"satellite/pkg"
	"strings"
)

// Run describe path to file
// https://docs.docker.com/compose/reference/exec/
type Exec struct {
	DockerCompose `yaml:",inline"`
	User          string   `yaml:"user"`
	Workdir       string   `yaml:"workdir"`
	Detach        bool     `yaml:"detach"`
	Env           []string `yaml:"environment"`
}

func (e *Exec) GetWorkdir() string {
	if e.Workdir != "" {
		return fmt.Sprintf("--workdir %s", e.Workdir)
	}

	return ""
}

func (e *Exec) GetEnv() string {
	var args []string
	for _, v := range e.Env {
		args = append(args, fmt.Sprintf("-e %s", v))
	}
	return strings.Join(args, " ")
}
func (e *Exec) GetUserId() string {
	if e.User != "" {
		return fmt.Sprintf("-u %s", e.User)
	}

	return ""
}
func (e *Exec) GetDetached() string {
	if e.Detach {
		return "-d"
	}
	return ""
}
func (e *Exec) ToCommand(args []string) []string {
	var arguments []string

	if len(args) >= 1 {
		arguments = args[0:]
	}

	bc := pkg.MergeSliceOfString([]string{
		e.GetPath(),
		e.GetProjectDirectory(),
		e.GetVerbose(),
		e.GetProjectName(),
		"exec",
		e.GetDetached(),
		e.GetUserId(),
		e.GetEnv(),
		e.GetWorkdir(),
	})

	return append(bc, pkg.DeleteEmpty(arguments)...)
}
