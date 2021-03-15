package entity

import (
	"github.com/mamau/starter/libs"
	"sync"
)

var gOnce sync.Once
var gInstance *Gulp

type Gulp struct {
	*Command
}

func GetGulp(image, homeDir string, args []string) *Gulp {
	gOnce.Do(func() {
		gInstance = &Gulp{
			Command: &Command{
				Image:   image,
				HomeDir: homeDir,
				Args:    args,
				Config:  libs.GetConfig().GetGulp(),
			},
		}
	})

	return gInstance
}

func (g *Gulp) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		g.Config.GetDns(),
		g.workDirVolume(),
		g.projectVolume(),
		{g.getImage()},
		{g.fullCommand()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}
