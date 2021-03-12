package services

import (
	"strings"
)

type Composer struct {
	//composer.Config `yaml:"config"`
	PreCommands []string `yaml:"pre-commands"`
}

//func (c *Composer) GetConfig() *composer.Config {
//	return &c.Config
//}

func (c *Composer) ToCommand() string {
	return strings.Join(c.PreCommands, "; ")
}
