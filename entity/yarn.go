package entity

import (
	"sync"

	"github.com/mamau/starter/libs"
)

var once sync.Once
var instance *Yarn

type Yarn struct {
	*Command
}

func NewYarn(version string, args []string) *Yarn {
	once.Do(func() {
		instance = &Yarn{
			Command: &Command{
				Name:    "yarn",
				Image:   "node",
				HomeDir: "/home/node",
				Version: version,
				Args:    args,
				Config:  libs.GetConfig().GetYarn(),
			},
		}
	})

	return instance
}
