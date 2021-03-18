package entity

import (
	"sync"

	"github.com/mamau/starter/libs"
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
				Image:   "mamau/bower",
				HomeDir: "/home/node",
				Args:    args,
				Config:  libs.GetConfig().GetBower(),
			},
		}
	})

	return bInstance
}
