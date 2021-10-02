package docker_compose

import (
	"fmt"
	"satellite/pkg"
)

// DockerCompose describe path to file
// https://docs.docker.com/compose/reference/
type DockerCompose struct {
	Name             string `yaml:"name" validate:"required,min=1"`
	Path             string `yaml:"path"`
	ProjectDirectory string `yaml:"project-directory"`
	ProjectName      string `yaml:"project-name"`
	LogLevel         string `yaml:"log-level"`
	Description      string `yaml:"description"`
	Verbose          bool   `yaml:"verbose"`
}

func (d *DockerCompose) GetExecCommand() string {
	return "docker-compose"
}

func (d *DockerCompose) GetDescription() string {
	return d.Description
}

func (d *DockerCompose) GetName() string {
	return d.Name
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
		d.GetVerbose(),
		d.GetProjectName(),
		command,
	})

	return append(bc, pkg.DeleteEmpty(arguments)...)
}
