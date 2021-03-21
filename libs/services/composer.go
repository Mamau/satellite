package services

import "github.com/mamau/starter/libs/services/composer"

type Composer struct {
	Docker          `yaml:",inline"`
	composer.Config `yaml:"config"`
}
