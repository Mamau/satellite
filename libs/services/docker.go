package services

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

func (d *Docker) GetPostCommands() string {
	return strings.Join(d.PostCommands, "; ")
}

func (d *Docker) GetCacheDir() string {
	return d.CacheDir
}

func (d *Docker) GetWorkDir() string {
	return d.WorkDir
}

func (d *Docker) GetUserId() []string {
	if id := d.UserId; id != "" {
		return []string{
			"-u",
			d.UserId,
		}
	}
	return nil
}

func (d *Docker) GetEnvironmentVariables() []string {
	var envVars []string
	for _, v := range d.EnvVars {
		envVars = append(envVars, fmt.Sprintf("-e %s", v))
	}
	return envVars
}

func (d *Docker) GetVersion() string {
	return d.Version
}

func (d *Docker) GetHosts() []string {
	var hosts []string
	for _, v := range d.AddHosts {
		hosts = append(hosts, fmt.Sprintf("--add-host=%s", v))
	}
	return hosts
}

func (d *Docker) GetPorts() []string {
	var ports []string
	for _, v := range d.Ports {
		ports = append(ports, fmt.Sprintf("-p %s", v))
	}
	return ports
}

func (d *Docker) GetVolumes() []string {
	var volumes []string
	for _, v := range d.Volumes {
		volumes = append(volumes, fmt.Sprintf("-v %s", v))
	}
	return volumes
}

func (d *Docker) GetDns() []string {
	var dns []string
	for _, v := range d.Dns {
		dns = append(dns, fmt.Sprintf("--dns=%s", v))
	}

	return dns
}
