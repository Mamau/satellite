package entity

import (
	"sync"

	"github.com/mamau/starter/libs"
)

var gOnce sync.Once
var gInstance *Gulp

type Gulp struct {
	*Command
}

func NewGulp(args []string) *Gulp {
	gOnce.Do(func() {
		gInstance = &Gulp{
			Command: &Command{
				Image:   "mamau/gulp",
				HomeDir: "/home/node",
				Args:    args,
				Config:  libs.GetConfig().GetGulp(),
			},
		}
	})

	return gInstance
}
