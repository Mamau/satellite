package entity

import (
	"fmt"
	"strings"
	"sync"

	"github.com/mamau/starter/libs"

	"github.com/mamau/starter/config/yarn"

	"github.com/mamau/starter/config/docker"

	"github.com/mamau/starter/config"
)

var once sync.Once
var instance *Yarn

type Yarn struct {
	Config *yarn.Yarn
	*Command
}

func NewYarn(version string, args []string) *Yarn {
	once.Do(func() {
		instance = &Yarn{
			Config: config.GetConfig().GetYarn(),
			Command: &Command{
				CmdName: "yarn",
				Image:   "node",
				HomeDir: "/home/node",
				Version: version,
				Args:    args,
			},
		}
	})

	return instance
}

func (y *Yarn) GetDockerConfig() *docker.Docker {
	return &y.Config.Docker
}

func (y *Yarn) GetCommandConfig() *Command {
	return y.Command
}

func (y *Yarn) GetClientSignature(cmd []string) []string {
	command := append(y.configToCommand(), cmd...)
	return []string{"/bin/bash", "-c", strings.Join(command, " ")}
}

func (y *Yarn) configToCommand() []string {
	if y.Config.Config == nil {
		return []string{}
	}
	configCommands := libs.DeleteEmpty(y.Config.GetAll())
	if len(configCommands) == 0 {
		return []string{}
	}

	for i, v := range configCommands {
		configCommands[i] = v + ";"
	}

	return libs.MergeSliceOfString(configCommands)
}

func (y *Yarn) GetImage() string {
	if v := y.Config.GetVersion(); v != "" {
		return fmt.Sprintf("%s:%s", y.Image, v)
	}
	return y.Command.GetImage()
}
