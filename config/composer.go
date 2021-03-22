package config

import "github.com/mamau/starter/config/composer"

type Composer struct {
	Docker          `yaml:",inline"`
	composer.Config `yaml:"config"`
}
