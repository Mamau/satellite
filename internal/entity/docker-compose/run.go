package docker_compose

import (
	"fmt"
	"satellite/pkg"
	"strings"
)

// Run describe path to file
// https://docs.docker.com/compose/reference/run/
type Run struct {
	DockerCompose `yaml:",inline"`
	User          string   `yaml:"user"`
	Entrypoint    string   `yaml:"entrypoint"`
	Env           []string `yaml:"environment"`
	Ports         []string `yaml:"ports"`
	Detach        bool     `yaml:"detach"`
	NoDeps        bool     `yaml:"no-deps"`
}

func (r *Run) GetPorts() string {
	var ports []string
	for _, v := range r.Ports {
		ports = append(ports, fmt.Sprintf("-p %s", v))
	}
	return strings.Join(ports, " ")
}

func (r *Run) GetEnv() string {
	var args []string
	for _, v := range r.Env {
		args = append(args, fmt.Sprintf("-e %s", v))
	}
	return strings.Join(args, " ")
}

func (r *Run) GetEntryPoint() string {
	if r.Entrypoint != "" {
		return fmt.Sprintf("--entrypoint %s", r.Entrypoint)
	}

	return ""
}

func (r *Run) GetUserId() string {
	if r.User != "" {
		return fmt.Sprintf("-u %s", r.User)
	}

	return ""
}

func (r *Run) GetNoDeps() string {
	if r.NoDeps {
		return "--no-deps"
	}
	return ""
}

func (r *Run) GetDetached() string {
	if r.Detach {
		return "-d"
	}
	return ""
}

func (r *Run) ToCommand(args []string) []string {
	var arguments []string

	if len(args) >= 1 {
		arguments = args[0:]
	}

	bc := pkg.MergeSliceOfString([]string{
		r.GetPath(),
		r.GetMultiPath(),
		r.GetProjectDirectory(),
		r.GetVerbose(),
		r.GetProjectName(),
		"run",
		r.GetDetached(),
		r.GetNoDeps(),
		r.GetUserId(),
		r.GetEntryPoint(),
		r.GetEnv(),
		r.GetPorts(),
	})

	return append(bc, pkg.DeleteEmpty(arguments)...)
}
