package docker_compose

import (
	"fmt"
	"satellite/pkg"
	"strings"
)

// Build describe path to file
// https://docs.docker.com/compose/reference/build/
type Build struct {
	DockerCompose `yaml:",inline"`
	Memory        string   `yaml:"memory"`
	Compress      bool     `yaml:"compress"`
	ForceRm       bool     `yaml:"force-rm"`
	NoCache       bool     `yaml:"no-cache"`
	NoRm          bool     `yaml:"no-rm"`
	Parallel      bool     `yaml:"parallel"`
	Pull          bool     `yaml:"pull"`
	Quiet         bool     `yaml:"quiet"`
	BuildArgs     []string `yaml:"build-arg"`
}

func (b *Build) GetQuiet() string {
	if b.Quiet {
		return "--quiet"
	}
	return ""
}
func (b *Build) GetPull() string {
	if b.Pull {
		return "--pull"
	}
	return ""
}
func (b *Build) GetParallel() string {
	if b.Parallel {
		return "--parallel"
	}
	return ""
}

func (b *Build) GetNoRm() string {
	if b.NoRm {
		return "--no-rm"
	}
	return ""
}

func (b *Build) GetNoCache() string {
	if b.NoCache {
		return "--no-cache"
	}
	return ""
}
func (b *Build) GetMemory() string {
	if b.Memory != "" {
		return fmt.Sprintf("--memory %s", b.Memory)
	}
	return ""
}
func (b *Build) GetForceRm() string {
	if b.ForceRm {
		return "--force-rm"
	}
	return ""
}
func (b *Build) GetCompress() string {
	if b.Compress {
		return "--compress"
	}
	return ""
}

func (b *Build) GetBuildArgs() string {
	var args []string
	for _, v := range b.BuildArgs {
		args = append(args, fmt.Sprintf("--build-arg %s", v))
	}
	return strings.Join(args, " ")
}

func (b *Build) ToCommand(args []string) []string {

	return pkg.MergeSliceOfString([]string{
		b.GetPath(),
		b.GetProjectDirectory(),
		b.GetVerbose(),
		b.GetProjectName(),
		"build",
		b.GetBuildArgs(),
		b.GetCompress(),
		b.GetForceRm(),
		b.GetMemory(),
		b.GetNoCache(),
		b.GetNoRm(),
		b.GetParallel(),
		b.GetPull(),
		b.GetQuiet(),
	})
}
