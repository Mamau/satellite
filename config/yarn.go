package config

import "github.com/mamau/starter/config/docker"

type Yarn struct {
	docker.Docker `yaml:",inline"`
}
