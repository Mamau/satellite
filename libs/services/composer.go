package services

import (
	"fmt"
	"strings"
)

type Composer struct {
	PreCommands []string `yaml:"pre-commands"`
	Dns         []string `yaml:"dns"`
}

func (c *Composer) ToCommand() string {
	return strings.Join(c.PreCommands, "; ")
}

func (c *Composer) GetDns() []string {
	var dns []string
	for _, v := range c.Dns {
		dns = append(dns, fmt.Sprintf("--dns=%s", v))
	}

	return dns
}
