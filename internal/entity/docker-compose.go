package entity

import (
	"fmt"
	"strings"

	"satellite/pkg"
)

// DockerCompose describe path to file
// https://docs.docker.com/compose/reference/
type DockerCompose struct {
	Name             string   `yaml:"name" validate:"required,min=1"`
	Path             string   `yaml:"path"`
	ProjectDirectory string   `yaml:"project-directory"`
	ProjectName      string   `yaml:"project-name"`
	LogLevel         string   `yaml:"log-level"`
	Description      string   `yaml:"description"`
	EnvFile          string   `yaml:"env-file"`
	User             string   `yaml:"user"`
	BuildArgs        []string `yaml:"build-arg"`
	Verbose          bool     `yaml:"verbose"`
	Detach           bool     `yaml:"detach"`
	RemoveOrphans    bool     `yaml:"remove-orphans"`
}

func (d *DockerCompose) GetBuildArgs() string {
	var args []string
	for _, v := range d.BuildArgs {
		args = append(args, fmt.Sprintf("--build-arg %s", v))
	}
	return strings.Join(args, " ")
}

func (d *DockerCompose) GetUserId() string {
	if d.User != "" {
		return fmt.Sprintf("-u %s", d.User)
	}

	return ""
}

func (d *DockerCompose) GetExecCommand() string {
	return string(DOCKER_COMPOSE)
}

func (d *DockerCompose) GetDescription() string {
	return d.Description
}

func (d *DockerCompose) GetName() string {
	return d.Name
}

func (d *DockerCompose) GetDetached() string {
	if d.Detach {
		return "-d"
	}
	return ""
}

func (d *DockerCompose) GetRemoveOrphans() string {
	if d.RemoveOrphans {
		return "--remove-orphans"
	}
	return ""
}

func (d *DockerCompose) GetEnvFile() string {
	if d.EnvFile != "" {
		return fmt.Sprintf("--env-file=%s", d.EnvFile)
	}

	return ""
}

func (d *DockerCompose) GetVerbose() string {
	if d.Verbose {
		return "--verbose"
	}
	return ""
}

func (d *DockerCompose) GetLogLevel() string {
	if d.LogLevel != "" {
		return fmt.Sprintf("--log-level %s", d.LogLevel)
	}
	return ""
}

func (d *DockerCompose) GetProjectName() string {
	if d.ProjectName != "" {
		return fmt.Sprintf("--project-name %s", d.ProjectName)
	}
	return ""
}

func (d *DockerCompose) GetPath() string {
	if d.Path != "" {
		return fmt.Sprintf("--file %s", d.Path)
	}
	return ""
}

func (d *DockerCompose) GetProjectDirectory() string {
	if d.ProjectDirectory != "" {
		return fmt.Sprintf("--project-directory %s", d.ProjectDirectory)
	}

	return ""
}

func (d *DockerCompose) ToCommand(args []string) []string {
	var command string
	var arguments []string

	if len(args) >= 1 {
		command = args[0]
	}
	if len(args) >= 2 {
		arguments = args[1:]
	}

	bc := pkg.MergeSliceOfString([]string{
		d.GetPath(),
		d.GetProjectDirectory(),
		d.GetUserId(),
		d.GetDetached(),
		d.GetEnvFile(),
		d.GetRemoveOrphans(),
		d.GetBuildArgs(),
		d.GetVerbose(),
		d.GetProjectName(),
		command,
	})
	configurator := newPureConfigConfigurator(bc, arguments)
	return append(bc, configurator.getClientCommand()...)
}
