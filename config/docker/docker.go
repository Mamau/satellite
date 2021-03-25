package docker

import (
	"fmt"
	"strings"
)

type Docker struct {
	Version      string   `yaml:"version"`
	UserId       string   `yaml:"user-id"`
	WorkDir      string   `yaml:"work-dir"`
	CacheDir     string   `yaml:"cache-dir"`
	PreCommands  []string `yaml:"pre-commands"`
	PostCommands []string `yaml:"post-commands"`
	Dns          []string `yaml:"dns"`
	Volumes      []string `yaml:"volumes"`
	Ports        []string `yaml:"ports"`
	AddHosts     []string `yaml:"add-hosts"`
	EnvVars      []string `yaml:"environment-variables"`
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

//func (d *Docker) workDir(dirname string) []string {
//	dir := dirname
//	if d.WorkDir != "" {
//		dir = d.WorkDir
//	}
//	if dir == "" {
//		return nil
//	}
//	return []string{
//		fmt.Sprintf("--workdir=%s", dirname),
//	}
//}

func (d *Docker) GetCacheVolume() string {
	if d.CacheDir == "" {
		return ""
	}

	return fmt.Sprintf("-v %s:/tmp", d.CacheDir)
}

func (d *Docker) GetWorkDir() string {
	if d.WorkDir == "" {
		return ""
	}

	return fmt.Sprintf("--workdir=%s", d.WorkDir)
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
