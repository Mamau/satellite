package services

import "strings"

type Gulp struct {
	PreCommands []string `yaml:"pre-commands"`
}

func (g *Gulp) ToCommand() string {
	return strings.Join(g.PreCommands, "; ")
}
