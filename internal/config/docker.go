package config

import (
	"satellite/internal/entity"
	"satellite/internal/entity/docker"
)

type Docker struct {
	Pulls []docker.Pull `yaml:"pull"`
	Runs  []docker.Run  `yaml:"run"`
	Execs []docker.Exec `yaml:"exec"`
}

func (d *Docker) GetTypes() []string {
	return []string{
		"pull",
		"run",
		"exec",
	}
}

func (d *Docker) GetCommands() []entity.Runner {
	var list []entity.Runner

	for _, v := range d.Runs {
		list = append(list, v)
	}
	for _, v := range d.Pulls {
		list = append(list, v)
	}
	for _, v := range d.Execs {
		list = append(list, v)
	}

	return list
}
