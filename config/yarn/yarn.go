package yarn

import "github.com/mamau/starter/config/docker"

type Yarn struct {
	docker.Docker `yaml:",inline"`
	*Config       `yaml:"config"`
}
