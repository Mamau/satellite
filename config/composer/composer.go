package composer

import (
	"github.com/mamau/starter/config/docker"
)

type Composer struct {
	docker.Docker `yaml:",inline"`
	*Config       `yaml:"config"`
}
