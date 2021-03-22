package config

import "github.com/mamau/starter/config/docker"

type Bower struct {
	docker.Docker `yaml:",inline"`
}
