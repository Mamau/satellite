package collector

import (
	"github.com/mamau/starter/config/docker"
)

type Collectorable interface {
	GetProjectVolume() string
	GetDockerConfig() *docker.Docker
	GetClientSignature(cmd []string) []string
	GetImage() string
	GetWorkDir() string
	GetImageCommand() string
	GetClientCommand() string
}
