package docker

import (
	"fmt"
	"os/user"
	"runtime"
	"strings"
)

type Docker struct {
	Name          string   `yaml:"name"`
	ContainerName string   `yaml:"container-name"`
	Image         string   `yaml:"image"`
	Command       string   `yaml:"command"`
	ImageCommand  string   `yaml:"image-command"`
	Flags         string   `yaml:"flags"`
	HomeDir       string   `yaml:"home-dir"`
	Version       string   `yaml:"version"`
	UserId        string   `yaml:"user-id"`
	WorkDir       string   `yaml:"work-dir"`
	Network       string   `yaml:"network"`
	Description   string   `yaml:"description"`
	Type          string   `yaml:"type"`
	Path          string   `yaml:"path"`
	ProjectName   string   `yaml:"project-name"`
	LogLevel      string   `yaml:"log-level"`
	Verbose       bool     `yaml:"verbose"`
	BinBash       bool     `yaml:"bin-bash"`
	SkipArgs      bool     `yaml:"skip-args"`
	Detach        bool     `yaml:"detach"`
	CleanUp       bool     `yaml:"clean-up"`
	PreCommands   []string `yaml:"pre-commands"`
	PostCommands  []string `yaml:"post-commands"`
	Dns           []string `yaml:"dns"`
	Volumes       []string `yaml:"volumes"`
	Ports         []string `yaml:"ports"`
	AddHosts      []string `yaml:"add-hosts"`
	EnvVars       []string `yaml:"environment-variables"`
}

func (d *Docker) GetVerbose() string {
	if d.Verbose {
		return "--verbose"
	}
	return ""
}

func (d *Docker) GetLogLevel() string {
	ll := LogLevel(d.LogLevel)
	if ll.IsAllowed() && d.LogLevel != "" {
		return fmt.Sprintf("--log-level %s", d.LogLevel)
	}

	return ""
}

func (d *Docker) GetProjectName() string {
	if d.ProjectName != "" {
		return fmt.Sprintf("--project-name %s", d.ProjectName)
	}
	return ""
}

func (d *Docker) GetPath() string {
	if d.Path != "" {
		return fmt.Sprintf("--file %s", d.Path)
	}
	return ""
}

func (d *Docker) GetType() Exec {
	dt := Exec(d.Type)
	if dt.IsAllowed() {
		return dt
	}

	return RUN
}

func (d *Docker) GetNetwork() string {
	if d.Network != "" {
		return fmt.Sprintf("--network %s", d.Network)
	}
	return ""
}

func (d *Docker) GetContainerName() string {
	if d.ContainerName != "" {
		return fmt.Sprintf("--name %s", d.ContainerName)
	}
	return ""
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
	if d.Detach {
		return "-d"
	}
	return ""
}

func (d *Docker) GetFlags() string {
	flags := d.Flags
	if flags == "" {
		flags = "-ti"
	}
	if runtime.GOOS == "windows" {
		flags = strings.Replace(strings.ToLower(flags), "-t", "-", -1)
		flags = strings.Replace(strings.ToLower(flags), "t", "", -1)
		if flags == "-" {
			flags = ""
		}
	}
	return flags
}

func (d *Docker) GetPreCommands() []string {
	if len(d.PreCommands) == 0 {
		return nil
	}

	commands := strings.Join(d.PreCommands, "; ")
	return strings.Split(commands, " ")
}

func (d *Docker) GetPostCommands() []string {
	if len(d.PostCommands) == 0 {
		return nil
	}

	commands := strings.Join(d.PostCommands, "; ")
	return strings.Split(commands, " ")
}

func (d *Docker) GetWorkDir() string {
	if d.WorkDir != "" {
		return fmt.Sprintf("--workdir=%s", d.WorkDir)
	}

	if d.HomeDir != "" {
		return fmt.Sprintf("--workdir=%s", d.HomeDir)
	}

	return ""
}

func (d *Docker) GetUserId() string {
	userID := d.UserId

	if userID == "" {
		return ""
	}

	if userID == "$(uid)" {
		currentUser, err := user.Current()
		if err != nil {
			return ""
		}
		userID = currentUser.Uid
	}

	return fmt.Sprintf("-u %s", userID)
}

func (d *Docker) GetEnvironmentVariables() string {
	var envVars []string
	for _, v := range d.EnvVars {
		envVars = append(envVars, fmt.Sprintf("-e %s", v))
	}
	return strings.Join(envVars, " ")
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

func (d *Docker) GetImage() string {
	imageName := d.Image
	if imageName == "" {
		imageName = d.Name
	}

	if d.Version != "" {
		return fmt.Sprintf("%s:%s", imageName, d.Version)
	}
	return imageName
}

func (d *Docker) GetImageCommand() string {
	if len(d.GetPreCommands()) > 0 || len(d.GetPostCommands()) > 0 {
		return "/bin/bash -c"
	}

	if d.BinBash == true {
		return "/bin/bash -c " + d.ImageCommand
	}

	return d.ImageCommand
}
