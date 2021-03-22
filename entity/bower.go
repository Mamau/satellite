package entity

import (
	"github.com/mamau/starter/config"
	"sync"
)

var bOnce sync.Once
var bInstance *Bower

type Bower struct {
	*Command
}

func NewBower(args []string) *Bower {
	bOnce.Do(func() {
		bInstance = &Bower{
			Command: &Command{
				Image:        "mamau/bower",
				HomeDir:      "/home/node",
				Args:         args,
				DockerConfig: config.GetConfig().GetBower(),
			},
		}
	})

	return bInstance
}

func (b *Bower) CollectCommand() []string {
	return append(b.dockerDataToCommand(), b.fullCommand())
}
