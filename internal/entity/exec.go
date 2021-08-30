package entity

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/mamau/satellite/pkg"
)

// Exec documentation for service params
// https://docs.docker.com/engine/reference/commandline/exec
type Exec struct {
	Name          string   `yaml:"name"`
	ContainerName string   `yaml:"container-name"`
	EnvFile       string   `yaml:"env-file"`
	User          string   `yaml:"user"`
	WorkDir       string   `yaml:"workdir"`
	Description   string   `yaml:"description"`
	Beginning     string   `yaml:"beginning"`
	Detach        bool     `yaml:"detach"`
	Interactive   bool     `yaml:"interactive"`
	Tty           bool     `yaml:"tty"`
	BinBash       bool     `yaml:"bin-bash"`
	PreCommands   []string `yaml:"pre-commands"`
	PostCommands  []string `yaml:"post-commands"`
	Env           []string `yaml:"env"`
}

func (e *Exec) GetDescription() string {
	return e.Description
}

func (e *Exec) GetName() string {
	return e.Name
}

func (e *Exec) GetExecCommand() string {
	return string(DOCKER)
}

func (e *Exec) ToCommand(args []string) []string {
	bc := pkg.MergeSliceOfString([]string{
		"exec",
		e.GetDetached(),
		e.GetFlags(),
		e.GetUserId(),
		e.GetEnv(),
		e.GetEnvFile(),
		e.GetWorkDir(),
		e.GetContainerName(),
	})
	args = append(e.GetStartCommand(), args...)
	configurator := newConfigConfigurator(bc, args, e)
	return append(bc, configurator.getClientCommand()...)
}

func (e *Exec) GetStartCommand() []string {
	if e.Beginning != "" {
		return strings.Split(e.Beginning, " ")
	}

	return []string{}
}

func (e *Exec) GetPreCommands() []string {
	if len(e.PreCommands) == 0 {
		return nil
	}

	commands := strings.Join(e.PreCommands, "; ")
	return strings.Split(commands, " ")
}

func (e *Exec) GetPostCommands() []string {
	if len(e.PostCommands) == 0 {
		return nil
	}

	commands := strings.Join(e.PostCommands, "; ")
	return strings.Split(commands, " ")
}

func (e *Exec) GetBinBash() bool {
	return e.BinBash
}

func (e *Exec) GetEnv() string {
	var envVars []string
	for _, v := range e.Env {
		envVars = append(envVars, fmt.Sprintf("--env %s", v))
	}
	return strings.Join(envVars, " ")
}

func (e *Exec) GetFlags() string {
	var flags []string
	var flagData string

	if e.Detach {
		return ""
	}

	if e.Interactive {
		flags = append(flags, "i")
	}

	if e.Tty && runtime.GOOS != "windows" {
		flags = append(flags, "t")
	}

	if len(flags) != 0 {
		flagData = "-" + strings.Join(flags, "")
	}

	return flagData
}

func (e *Exec) GetDetached() string {
	if e.Detach {
		return "-d"
	}
	return ""
}

func (e *Exec) GetWorkDir() string {
	if e.WorkDir != "" {
		return fmt.Sprintf("--workdir=%s", e.WorkDir)
	}

	return ""
}

func (e *Exec) GetUserId() string {
	if e.User != "" {
		return fmt.Sprintf("--user %s", e.User)
	}

	return ""
}

func (e *Exec) GetEnvFile() string {
	if e.EnvFile != "" {
		return fmt.Sprintf("--env-file %s", e.EnvFile)
	}

	return ""
}

func (e *Exec) GetContainerName() string {
	return e.ContainerName
}
