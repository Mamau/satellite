package services

import "github.com/mamau/starter/libs/services/composer"

type Composer struct {
	composer.Config `yaml:"config"`
}

func (c *Composer) GetConfig() *composer.Config {
	return &c.Config
}
