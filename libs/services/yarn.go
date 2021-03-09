package services

import (
	"strings"
)

type Yarn struct {
	PreCommands []string `yaml:"pre-commands"`
}

func (y *Yarn) ToCommand() string {
	return strings.Join(y.PreCommands, "; ")
}
