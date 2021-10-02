package docker_compose

import (
	"fmt"
	"satellite/pkg"
)

// Down describe path to file
// https://docs.docker.com/compose/reference/down/
type Down struct {
	DockerCompose `yaml:",inline"`
	Rmi           string `yaml:"rmi"`
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

	return pkg.MergeSliceOfString([]string{
		d.GetPath(),
		d.GetProjectDirectory(),
		d.GetVerbose(),
		d.GetProjectName(),
		"down",
		d.GetRemoveOrphans(),
		d.GetRmi(),
		d.GetRemoveVolumes(),
	})
}
