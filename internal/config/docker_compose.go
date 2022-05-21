package config

import (
	"satellite/internal/entity"
	docker_compose "satellite/internal/entity/docker-compose"
)

type DockerCompose struct {
	Run   []docker_compose.Run   `yaml:"run"`
	Up    []docker_compose.Up    `yaml:"up"`
	Down  []docker_compose.Down  `yaml:"down"`
	Exec  []docker_compose.Exec  `yaml:"exec"`
	Build []docker_compose.Build `yaml:"build"`
}

func (d *DockerCompose) GetTypes() []string {
	return []string{
		"run",
		"up",
		"down",
		"exec",
		"build",
	}
}

func (d *DockerCompose) GetCommands() []entity.Runner {
	var list []entity.Runner

	for _, v := range d.Run {
		list = append(list, &v)
	}
	for _, v := range d.Up {
		list = append(list, &v)
	}
	for _, v := range d.Down {
		list = append(list, &v)
	}
	for _, v := range d.Exec {
		list = append(list, &v)
	}
	for _, v := range d.Build {
		list = append(list, &v)
	}
	return list
}
