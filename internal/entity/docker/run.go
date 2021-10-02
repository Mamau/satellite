package docker

import (
	"fmt"
	"os/exec"
	"runtime"
	"satellite/internal/entity"
	"strings"

	"github.com/gookit/color"

	"satellite/pkg"
)

// Run documentation for service params
// https://docs.docker.com/engine/reference/commandline/run
type Run struct {
	docker        `yaml:",inline"`
	ContainerName string   `yaml:"container-name"`
	Image         string   `yaml:"image" validate:"required,min=1"`
	Version       string   `yaml:"version"`
	WorkDir       string   `yaml:"workdir"`
	Network       string   `yaml:"network"`
	Hostname      string   `yaml:"hostname"`
	EnvFile       string   `yaml:"env-file"`
	User          string   `yaml:"user"`
	Beginning     string   `yaml:"beginning"`
	Detach        bool     `yaml:"detach"`
	Interactive   bool     `yaml:"interactive"`
	Tty           bool     `yaml:"tty"`
	CleanUp       bool     `yaml:"clean-up"`
	BinBash       bool     `yaml:"bin-bash"`
	PreCommands   []string `yaml:"pre-commands"`
	PostCommands  []string `yaml:"post-commands"`
	Dns           []string `yaml:"dns"`
	Volumes       []string `yaml:"volumes"`
	Ports         []string `yaml:"ports"`
	AddHosts      []string `yaml:"add-hosts"`
	Env           []string `yaml:"env"`
}

func (r *Run) GetDetached() string {
	if r.Detach {
		return "-d"
	}
	return ""
}

func (r *Run) GetEnvFile() string {
	if r.EnvFile != "" {
		return fmt.Sprintf("--env-file %s", r.EnvFile)
	}

	return ""
}

func (r *Run) GetHostname() string {
	if r.Hostname != "" {
		return fmt.Sprintf("--hostname %s", r.Hostname)
	}

	return ""
}

func (r *Run) GetWorkDir() string {
	if r.WorkDir != "" {
		return fmt.Sprintf("--workdir=%s", r.WorkDir)
	}

	return ""
}

func (r *Run) GetFlags() string {
	var flags []string
	var flagData string

	if r.Detach {
		return ""
	}

	if r.Interactive {
		flags = append(flags, "i")
	}

	if r.Tty && runtime.GOOS != "windows" {
		flags = append(flags, "t")
	}

	if len(flags) != 0 {
		flagData = "-" + strings.Join(flags, "")
	}

	return flagData
}

func (r *Run) GetCleanUp() string {
	if r.CleanUp {
		return "--rm"
	}
	return ""
}

func (r *Run) GetNetwork() string {
	if r.Network != "" {
		return fmt.Sprintf("--network %s", r.Network)
	}
	return ""
}

func (r *Run) GetUserId() string {
	if r.User != "" {
		return fmt.Sprintf("--user %s", r.User)
	}

	return ""
}

func (r *Run) GetEnv() string {
	var envVars []string
	for _, v := range r.Env {
		envVars = append(envVars, fmt.Sprintf("--env %s", v))
	}
	return strings.Join(envVars, " ")
}

func (r *Run) GetHosts() string {
	var hosts []string
	for _, v := range r.AddHosts {
		hosts = append(hosts, fmt.Sprintf("--add-host=%s", v))
	}
	return strings.Join(hosts, " ")
}

func (r *Run) GetPorts() string {
	var ports []string
	for _, v := range r.Ports {
		ports = append(ports, fmt.Sprintf("-p %s", v))
	}
	return strings.Join(ports, " ")
}

func (r *Run) GetDns() string {
	var dns []string
	for _, v := range r.Dns {
		dns = append(dns, fmt.Sprintf("--dns=%s", v))
	}

	return strings.Join(dns, " ")
}

func (r *Run) GetVolumes() string {
	var volumes []string
	for _, v := range r.Volumes {
		volumes = append(volumes, fmt.Sprintf("-v %s", v))
	}
	return strings.Join(volumes, " ")
}

func (r *Run) GetContainerName() string {
	if r.ContainerName != "" {
		return fmt.Sprintf("--name %s", r.ContainerName)
	}

	if r.Name != "" {
		return fmt.Sprintf("--name %s", r.Name)
	}
	return ""
}

func (r *Run) GetImage() string {
	if r.Version != "" {
		return fmt.Sprintf("%s:%s", r.Image, r.Version)
	}
	return r.Image
}

func (r *Run) GetPreCommands() []string {
	if len(r.PreCommands) == 0 {
		return nil
	}

	commands := strings.Join(r.PreCommands, "; ")
	return strings.Split(commands, " ")
}

func (r *Run) GetPostCommands() []string {
	if len(r.PostCommands) == 0 {
		return nil
	}

	commands := strings.Join(r.PostCommands, "; ")
	return strings.Split(commands, " ")
}

func (r *Run) GetBinBash() bool {
	return r.BinBash
}

func (r *Run) GetExecCommand() string {
	return string(entity.DOCKER)
}

func (r *Run) ToCommand(args []string) []string {
	r.createNetwork()

	bc := pkg.MergeSliceOfString([]string{
		"run",
		r.GetDetached(),
		r.GetFlags(),
		r.GetCleanUp(),
		r.GetHostname(),
		r.GetNetwork(),
		r.GetUserId(),
		r.GetEnv(),
		r.GetEnvFile(),
		r.GetHosts(),
		r.GetPorts(),
		r.GetDns(),
		r.GetWorkDir(),
		r.GetVolumes(),
		r.GetContainerName(),
		r.GetImage(),
	})
	args = append(r.GetBeginning(), args...)
	configurator := NewConfigConfigurator(bc, args, r)
	return append(bc, configurator.GetClientCommand()...)
}

func (r *Run) GetBeginning() []string {
	if r.Beginning != "" {
		return strings.Split(r.Beginning, " ")
	}

	return []string{}
}

func (r *Run) createNetwork() {
	if r.Network == "" {
		return
	}

	data := []string{
		"network",
		"inspect",
		r.Network,
	}

	cmd := exec.Command(r.GetExecCommand(), data...)
	network, _ := cmd.Output()

	if strings.TrimSuffix(string(network), "\n") == "[]" {
		color.Yellow.Printf("Creating network %s\n", r.Network)
		data := []string{
			"network",
			"create",
			r.Network,
		}
		cmd := exec.Command(r.GetExecCommand(), data...)
		out, _ := cmd.Output()

		color.Cyan.Printf("%s\n", out)
	}
}
