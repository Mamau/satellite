package entity

import (
	"fmt"

	"github.com/mamau/satellite/pkg"
)

// DockerCompose describe path to file
// https://docs.docker.com/compose/reference/
type DockerCompose struct {
	Name             string `yaml:"name"`
	Path             string `yaml:"path"`
	ProjectDirectory string `yaml:"project-directory"`
	ProjectName      string `yaml:"project-name"`
	LogLevel         string `yaml:"log-level"`
	Command          string `yaml:"command"`
	Description      string `yaml:"description"`
	Verbose          bool   `yaml:"verbose"`
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

func (d *DockerCompose) GetCommand() string {
	if d.Command != "" {
		return d.Command
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
	bc := pkg.MergeSliceOfString([]string{
		d.GetCommand(),
		d.GetPath(),
		d.GetProjectDirectory(),
		d.GetVerbose(),
		d.GetProjectName(),
	})
	configurator := newPureConfigConfigurator(bc, args)
	return append(bc, configurator.getClientCommand()...)
}
