package services

import (
	"fmt"
	"strings"
)

type Gulp struct {
	PreCommands []string `yaml:"pre-commands"`
	Dns         []string `yaml:"dns"`
}

func (g *Gulp) ToCommand() string {
	return strings.Join(g.PreCommands, "; ")
}

func (g *Gulp) GetDns() []string {
	var dns []string
	for _, v := range g.Dns {
		dns = append(dns, fmt.Sprintf("--dns=%s", v))
	}

	return dns
}
