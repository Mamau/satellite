package services

import (
	"fmt"
	"strings"
)

type Yarn struct {
	PreCommands []string `yaml:"pre-commands"`
	Dns         []string `yaml:"dns"`
}

func (y *Yarn) ToCommand() string {
	return strings.Join(y.PreCommands, "; ")
}

func (y *Yarn) GetDns() []string {
	var dns []string
	for _, v := range y.Dns {
		dns = append(dns, fmt.Sprintf("--dns=%s", v))
	}

	return dns
}
