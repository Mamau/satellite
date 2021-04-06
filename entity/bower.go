package entity

import (
	"github.com/mamau/starter/config/docker"

	"github.com/mamau/starter/config"
)

type Bower struct {
	Config *config.Bower
	*Command
}

func NewBower(args []string) *Bower {
	return &Bower{
		Config: config.GetConfig().GetBower(),
		Command: &Command{
			CmdName: "bower",
			Image:   "mamau/bower",
			HomeDir: "/home/node",
			Args:    args,
		},
	}
}

//func (b *Bower) GetWorkDir() string {
//
//}

func (b *Bower) GetDockerConfig() *docker.Docker {
	if b.Config == nil {
		return nil
	}
	return &b.Config.Docker
}

func (b *Bower) GetCommandConfig() *Command {
	return b.Command
}

func (b *Bower) GetClientSignature(cmd []string) []string {
	return cmd
}

func (b *Bower) GetImage() string {
	return b.Image
}
