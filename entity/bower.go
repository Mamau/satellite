package entity

import (
	"github.com/mamau/starter/libs"
	"sync"
)

var bOnce sync.Once
var bInstance *Bower

type Bower struct {
	*Command
}

func GetBower(image, homeDir string, args []string) *Bower {
	bOnce.Do(func() {
		bInstance = &Bower{
			Command: &Command{
				Image:   image,
				HomeDir: homeDir,
				Args:    args,
				Config:  libs.GetConfig().GetBower(),
			},
		}
	})

	return bInstance
}
func (b *Bower) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		b.Config.GetDns(),
		b.workDirVolume(),
		b.projectVolume(),
		{b.getImage()},
		{b.fullCommand()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}
