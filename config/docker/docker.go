package docker

import (
	"fmt"
	"strings"
)

type Docker struct {
	Name           string   `yaml:"name"`
	Image          string   `yaml:"image"`
	Command        string   `yaml:"command"`
	ExecuteCommand string   `yaml:"execute-command"`
	Flags          string   `yaml:"flags"`
	HomeDir        string   `yaml:"home-dir"`
	Version        string   `yaml:"version"`
	UserId         string   `yaml:"user-id"`
	WorkDir        string   `yaml:"work-dir"`
	Detach         bool     `yaml:"detach"`
	CleanUp        bool     `yaml:"clean-up"`
	PreCommands    []string `yaml:"pre-commands"`
	PostCommands   []string `yaml:"post-commands"`
	Dns            []string `yaml:"dns"`
	Volumes        []string `yaml:"volumes"`
	Ports          []string `yaml:"ports"`
	AddHosts       []string `yaml:"add-hosts"`
	EnvVars        []string `yaml:"environment-variables"`
}

func (d *Docker) GetImage() string {
	if d.Image != "" {
		return d.Image
	}
	return ""
}

func (d *Docker) GetExecCommand() string {
	if d.ExecuteCommand != "" {
		return d.ExecuteCommand
	}
	return "/bin/bash -c"
}

func (d *Docker) GetCleanUp() string {
	if d.CleanUp {
		return "--rm"
	}
	return ""
}

func (d *Docker) GetDockerCommand() string {
	if d.Command != "" {
		return d.Command
	}
	return "run"
}

func (d *Docker) GetDetached() string {
	if d.GetDockerCommand() == "pull" {
		return ""
	}

	if d.Detach {
		return "-d"
	}
	return ""
}

func (d *Docker) GetFlags() string {
	if d.GetDockerCommand() == "pull" || d.Detach {
		return ""
	}

	if d.Flags != "" {
		return d.Flags
	}
	return "-ti"
}

func (d *Docker) GetPreCommands() string {
	return strings.Join(d.PreCommands, "; ")
}

func (d *Docker) SetPreCommands(pc []string) {
	d.PreCommands = pc
}

func (d *Docker) SetPostCommands(pc []string) {
	d.PostCommands = pc
}

func (d *Docker) GetPostCommands() string {
	return strings.Join(d.PostCommands, "; ")
}

func (d *Docker) GetWorkDir() string {
	if d.WorkDir != "" {
		return fmt.Sprintf("--workdir=%s", d.WorkDir)
	}

	return ""
}

func (d *Docker) SetWorkDir(wd string) {
	d.WorkDir = wd
}

func (d *Docker) GetUserId() string {
	if d.UserId == "" {
		return ""
	}

	return fmt.Sprintf("-u %s", d.UserId)
}

func (d *Docker) GetEnvironmentVariables() string {
	var envVars []string
	for _, v := range d.EnvVars {
		envVars = append(envVars, fmt.Sprintf("-e %s", v))
	}
	return strings.Join(envVars, " ")
}

func (d *Docker) SetVersion(v string) {
	d.Version = v
}

func (d *Docker) GetVersion() string {
	return d.Version
}

func (d *Docker) GetHosts() string {
	var hosts []string
	for _, v := range d.AddHosts {
		hosts = append(hosts, fmt.Sprintf("--add-host=%s", v))
	}
	return strings.Join(hosts, " ")
}

func (d *Docker) GetPorts() string {
	var ports []string
	for _, v := range d.Ports {
		ports = append(ports, fmt.Sprintf("-p %s", v))
	}
	return strings.Join(ports, " ")
}

func (d *Docker) GetVolumes() string {
	var volumes []string
	for _, v := range d.Volumes {
		volumes = append(volumes, fmt.Sprintf("-v %s", v))
	}
	return strings.Join(volumes, " ")
}

func (d *Docker) GetDns() string {
	var dns []string
	for _, v := range d.Dns {
		dns = append(dns, fmt.Sprintf("--dns=%s", v))
	}

	return strings.Join(dns, " ")
}
