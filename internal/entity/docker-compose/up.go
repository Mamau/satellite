package docker_compose

import "satellite/pkg"

// Run describe path to file
// https://docs.docker.com/compose/reference/up/
type Up struct {
	DockerCompose `yaml:",inline"`
	Detach        bool `yaml:"detach"`
	NoDeps        bool `yaml:"no-deps"`
	Build         bool `yaml:"build"`
	NoBuild       bool `yaml:"no-build"`
	NoStart       bool `yaml:"no-start"`
	RemoveOrphans bool `yaml:"remove-orphans"`
}

func (u *Up) GetBuild() string {
	if u.Build {
		return "--build"
	}
	return ""
}

func (u *Up) GetNoBuild() string {
	if u.NoBuild {
		return "--no-build"
	}
	return ""
}

func (u *Up) GetRemoveOrphans() string {
	if u.RemoveOrphans {
		return "--remove-orphans"
	}
	return ""
}

func (u *Up) GetNoStart() string {
	if u.NoStart {
		return "--no-start"
	}
	return ""
}

func (u *Up) GetDetached() string {
	if u.Detach {
		return "-d"
	}
	return ""
}

func (u *Up) GetNoDeps() string {
	if u.NoDeps {
		return "--no-deps"
	}
	return ""
}

func (u *Up) ToCommand(args []string) []string {

	return pkg.MergeSliceOfString([]string{
		u.GetPath(),
		u.GetProjectDirectory(),
		u.GetVerbose(),
		u.GetProjectName(),
		"up",
		u.GetDetached(),
		u.GetNoDeps(),
		u.GetNoStart(),
		u.GetRemoveOrphans(),
		u.GetNoBuild(),
		u.GetBuild(),
	})
}
