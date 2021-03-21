package composer

import (
	"fmt"
	"strings"
)

type Config struct {
	Repositories       []string `yaml:"repositories"`
	ProcessTimeout     string   `yaml:"process-timeout"`
	OptimizeAutoloader string   `yaml:"optimize-autoloader"`
}

func (c *Config) GetProcessTimeoutAsCommand() string {
	if c.ProcessTimeout == "" {
		return ""
	}
	cmd := fmt.Sprintf("process-timeout %s", c.ProcessTimeout)
	return c.valueToCommand(cmd)
}

func (c *Config) GetRepoAsCommand() string {
	if c.Repositories == nil {
		return ""
	}

	return c.listToCommand(c.Repositories)
}

func (c *Config) GetOptimizeAutoloaderAsCommand() string {
	if c.OptimizeAutoloader == "" {
		return ""
	}

	cmd := fmt.Sprintf("optimize-autoloader %s", c.OptimizeAutoloader)
	return c.valueToCommand(cmd)
}

func (c *Config) GetAll() []string {
	return []string{
		c.GetProcessTimeoutAsCommand(),
		c.GetRepoAsCommand(),
		c.GetOptimizeAutoloaderAsCommand(),
	}

	//return libs.DeleteEmpty(list)
}

func (c *Config) listToCommand(list []string) string {
	var lc []string
	for _, v := range list {
		lc = append(lc, c.valueToCommand(v))
	}

	return strings.Join(lc, "; ")
}

func (c *Config) valueToCommand(v string) string {
	return fmt.Sprintf("composer config --global %s", v)
}
