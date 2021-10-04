package docker_compose

import (
	"fmt"
	"satellite/pkg"
)

// Down describe path to file
// https://docs.docker.com/compose/reference/down/
type Down struct {
	DockerCompose `yaml:",inline"`
	Rmi           string `yaml:"rmi" validate:"in=all local"`
	RemoveVolumes bool   `yaml:"volumes"`
	RemoveOrphans bool   `yaml:"remove-orphans"`
}

func (d *Down) GetRemoveVolumes() string {
	if d.RemoveVolumes {
		return "--volumes"
	}
	return ""
}
func (d *Down) GetRmi() string {
	if d.Rmi != "" {
		return fmt.Sprintf("--rmi %s", d.Rmi)
	}
	return ""
}
func (d *Down) GetRemoveOrphans() string {
	if d.RemoveOrphans {
		return "--remove-orphans"
	}
	return ""
}
func (d *Down) ToCommand(args []string) []string {
	var arguments []string

	if len(args) >= 1 {
		arguments = args[0:]
	}

	bc := pkg.MergeSliceOfString([]string{
		d.GetPath(),
		d.GetProjectDirectory(),
		d.GetVerbose(),
		d.GetProjectName(),
		"down",
		d.GetRemoveOrphans(),
		d.GetRmi(),
		d.GetRemoveVolumes(),
	})

	return append(bc, pkg.DeleteEmpty(arguments)...)
}
