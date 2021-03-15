package services

import (
	"fmt"
	"strings"
)

type Bower struct {
	PreCommands []string `yaml:"pre-commands"`
	Dns         []string `yaml:"dns"`
}

func (b *Bower) ToCommand() string {
	return strings.Join(b.PreCommands, "; ")
}

func (b *Bower) GetDns() []string {
	var dns []string
	for _, v := range b.Dns {
		dns = append(dns, fmt.Sprintf("--dns=%s", v))
	}

	return dns
}
