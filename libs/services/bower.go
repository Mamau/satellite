package services

import "strings"

type Bower struct {
	PreCommands []string `yaml:"pre-commands"`
}

func (b *Bower) ToCommand() string {
	return strings.Join(b.PreCommands, "; ")
}
