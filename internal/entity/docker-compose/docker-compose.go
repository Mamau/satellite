package docker_compose

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"os/exec"
	"satellite/internal/entity"
	"satellite/pkg"
	"strings"
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
	MultiPath        []string `yaml:"multi-path"`
	Verbose          bool     `yaml:"verbose"`
}

func (d *DockerCompose) GetExecCommand() string {
	_, err := exec.LookPath(string(entity.DOCKER_COMPOSE))
	if err == nil {
		return string(entity.DOCKER_COMPOSE)
	}
	color.Warn.Println("You have no docker-compose. Checking for docker compose 2nd version...")

	cmd := exec.Command("docker", "compose")
	if err := cmd.Run(); err != nil {
		color.Red.Println("Oops... you need to install docker compose")
		os.Exit(1)
	}

	return string(entity.DOCKER_COMPOSE_2)
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

func (d *DockerCompose) GetMultiPath() string {
	var args []string
	for _, v := range d.MultiPath {
		args = append(args, fmt.Sprintf("--file %s", v))
	}
	return strings.Join(args, " ")
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
		d.GetMultiPath(),
		d.GetProjectDirectory(),
		d.GetVerbose(),
		d.GetProjectName(),
		command,
	})

	return append(bc, pkg.DeleteEmpty(arguments)...)
}
