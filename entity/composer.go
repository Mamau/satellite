package entity

import (
	"sync"

	"github.com/mamau/starter/libs"
)

var cOnce sync.Once
var cInstance *Composer

type Composer struct {
	*Command
}

func NewComposer(version string, args []string) *Composer {
	cOnce.Do(func() {
		cInstance = &Composer{
			Command: &Command{
				Name:    "composer",
				Image:   "composer",
				HomeDir: "/home/www-data",
				Version: version,
				Args:    args,
				Config:  libs.GetConfig().GetComposer(),
			},
		}
	})

	return cInstance
}
