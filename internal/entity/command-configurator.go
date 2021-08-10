package entity

import (
	"strings"

	"github.com/mamau/satellite/pkg"
)

type BinBasher interface {
	GetPreCommands() []string
	GetPostCommands() []string
	GetBinBash() bool
}

type commandConfigurator struct {
	configCmd []string
	preCmd    []string
	postCmd   []string
	mainCmd   []string
	isBinBash bool
}

func newPureConfigConfigurator(configCmd, mainCmd []string) *commandConfigurator {
	return &commandConfigurator{
		configCmd: configCmd,
		mainCmd:   mainCmd,
	}
}

func newConfigConfigurator(configCmd, mainCmd []string, binBasher BinBasher) *commandConfigurator {
	var bb bool
	if binBasher.GetBinBash() || len(binBasher.GetPreCommands()) > 0 || len(binBasher.GetPostCommands()) > 0 {
		bb = true
	}

	return &commandConfigurator{
		configCmd: configCmd,
		preCmd:    binBasher.GetPreCommands(),
		postCmd:   binBasher.GetPostCommands(),
		mainCmd:   mainCmd,
		isBinBash: bb,
	}
}

func (c *commandConfigurator) getClientCommand() []string {
	fullCmd := c.prepareCommand()

	if c.isBinBash {
		return append(c.binBash(), strings.Join(fullCmd, " "))
	}
	return append(c.binBash(), fullCmd...)
}

func (c *commandConfigurator) prepareCommand() []string {
	if len(c.preCmd) > 0 {
		c.preCmd[len(c.preCmd)-1] += ";"
	}

	if len(c.postCmd) > 0 {
		c.mainCmd[len(c.mainCmd)-1] += ";"
	}

	startCmd := append(c.preCmd, c.mainCmd...)
	fullCmd := append(startCmd, c.postCmd...)

	return pkg.DeleteEmpty(fullCmd)
}

func (c *commandConfigurator) binBash() []string {
	if c.isBinBash {
		return []string{"/bin/bash", "-c"}
	}

	return []string{}
}
