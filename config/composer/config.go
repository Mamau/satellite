package composer

import (
	"fmt"
	"strings"
)

type Config struct {
	Repositories       []string `yaml:"repositories"`
	ProcessTimeout     string   `yaml:"process-timeout"`
	OptimizeAutoloader string   `yaml:"optimize-autoloader"`
	GithubOauth        string   `yaml:"github-oauth"`
	GitlabOauth        string   `yaml:"gitlab-oauth"`
	GitlabToken        string   `yaml:"gitlab-token"`
	Lock               string   `yaml:"lock"`
}

func (c *Config) GetLock() string {
	if c.Lock == "" {
		return ""
	}
	cmd := fmt.Sprintf("lock %s", c.Lock)
	return c.valueToCommand(cmd)
}

func (c *Config) GetGitlabToken() string {
	if c.GitlabToken == "" {
		return ""
	}
	cmd := fmt.Sprintf("gitlab-token %s", c.GitlabToken)
	return c.valueToCommand(cmd)
}

func (c *Config) GetGitlabOauth() string {
	if c.GitlabOauth == "" {
		return ""
	}
	cmd := fmt.Sprintf("gitlab-oauth %s", c.GitlabOauth)
	return c.valueToCommand(cmd)
}

func (c *Config) GetGithubOauth() string {
	if c.GithubOauth == "" {
		return ""
	}
	cmd := fmt.Sprintf("github-oauth %s", c.GithubOauth)
	return c.valueToCommand(cmd)
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
		c.GetGithubOauth(),
		c.GetGitlabOauth(),
		c.GetGitlabToken(),
		c.GetLock(),
	}
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
