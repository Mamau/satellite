package entity

import (
	"github.com/mamau/starter/libs"
	"strings"
	"sync"
)

var once sync.Once
var instance *Yarn

type Yarn struct {
	*Command
}

func GetYarn(image, homeDir, version string, args []string) *Yarn {
	once.Do(func() {
		instance = &Yarn{
			Command: &Command{
				Image:   image,
				HomeDir: homeDir,
				Version: version,
				Args:    args,
				Config:  libs.GetConfig().GetYarn(),
			},
		}
	})

	return instance
}

func (y *Yarn) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		y.Config.GetDns(),
		y.workDirVolume(),
		y.projectVolume(),
		{y.getImage()},
		{"/bin/bash", "-c", y.fullCommand()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}

func (y *Yarn) fullCommand() string {
	return y.getConfigCommand() + y.getMainCommand() + y.getPostCommands()
}

func (y *Yarn) getMainCommand() string {
	mainCommand := append([]string{"yarn"}, y.Args...)
	return strings.Join(mainCommand, " ")
}
