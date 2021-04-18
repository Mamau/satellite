package collector

import (
	"github.com/mamau/starter/config/docker"
)

type Collectorable interface {
	GetDockerConfig() *docker.Docker
	GetImage() string
	GetWorkDir() string
	GetImageCommand() string
	GetClientCommand() string
}
