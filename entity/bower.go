package entity

import (
	"sync"

	"github.com/mamau/starter/config/docker"

	"github.com/mamau/starter/config"
)

var bOnce sync.Once
var bInstance *Bower

type Bower struct {
	Config *config.Bower
	*Command
}

func NewBower(args []string) *Bower {
	bOnce.Do(func() {
		bInstance = &Bower{
			Config: config.GetConfig().GetBower(),
			Command: &Command{
				CmdName: "bower",
				Image:   "mamau/bower",
				HomeDir: "/home/node",
				Args:    args,
			},
		}
	})

	return bInstance
}

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
