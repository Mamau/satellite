package yarn

import "fmt"

type Config struct {
	StrictSsl          string `yaml:"strict-ssl"`
	VersionTagPrefix   string `yaml:"version-tag-prefix"`
	VersionGitTag      string `yaml:"version-git-tag"`
	VersionCommitHooks string `yaml:"version-commit-hooks"`
	VersionGitSign     string `yaml:"version-git-sign"`
	BinLinks           string `yaml:"bin-links"`
	IgnoreScripts      string `yaml:"ignore-scripts"`
	IgnoreOptional     string `yaml:"ignore-optional"`
}

func (c *Config) GetAll() []string {
	return []string{
		c.GetStrictSsl(),
		c.GetVersionTagPrefix(),
		c.GetVersionGitTag(),
		c.GetVersionCommitHooks(),
		c.GetVersionGitSign(),
		c.GetBinLinks(),
		c.GetIgnoreScripts(),
		c.GetIgnoreOptional(),
	}
}

func (c *Config) GetIgnoreOptional() string {
	if c.IgnoreOptional == "" {
		return ""
	}

	cmd := fmt.Sprintf("ignore-optional %s", c.IgnoreOptional)
	return c.valueToCommand(cmd)
}

func (c *Config) GetIgnoreScripts() string {
	if c.IgnoreScripts == "" {
		return ""
	}

	cmd := fmt.Sprintf("ignore-scripts %s", c.IgnoreScripts)
	return c.valueToCommand(cmd)
}

func (c *Config) GetBinLinks() string {
	if c.BinLinks == "" {
		return ""
	}

	cmd := fmt.Sprintf("bin-links %s", c.BinLinks)
	return c.valueToCommand(cmd)
}

func (c *Config) GetVersionGitSign() string {
	if c.VersionGitSign == "" {
		return ""
	}

	cmd := fmt.Sprintf("version-git-sign %s", c.VersionGitSign)
	return c.valueToCommand(cmd)
}

func (c *Config) GetVersionCommitHooks() string {
	if c.VersionCommitHooks == "" {
		return ""
	}

	cmd := fmt.Sprintf("version-commit-hooks %s", c.VersionCommitHooks)
	return c.valueToCommand(cmd)
}

func (c *Config) GetVersionGitTag() string {
	if c.VersionGitTag == "" {
		return ""
	}

	cmd := fmt.Sprintf("version-git-tag %s", c.VersionGitTag)
	return c.valueToCommand(cmd)
}

func (c *Config) GetVersionTagPrefix() string {
	if c.VersionTagPrefix == "" {
		return ""
	}

	cmd := fmt.Sprintf("version-tag-prefix %s", c.VersionTagPrefix)
	return c.valueToCommand(cmd)
}

func (c *Config) GetStrictSsl() string {
	if c.StrictSsl == "" {
		return ""
	}

	cmd := fmt.Sprintf("strict-ssl %s", c.StrictSsl)
	return c.valueToCommand(cmd)
}

func (c *Config) valueToCommand(v string) string {
	return fmt.Sprintf("yarn config set %s --global", v)
}
