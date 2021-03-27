package entity

import (
	"fmt"
	"strings"
	"sync"

	"github.com/mamau/starter/config/docker"

	"github.com/mamau/starter/config"
)

var once sync.Once
var instance *Yarn

type Yarn struct {
	Config *config.Yarn
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
	return []string{"/bin/bash", "-c", strings.Join(cmd, " ")}
}

func (y *Yarn) GetImage() string {
	if v := y.Config.GetVersion(); v != "" {
		return fmt.Sprintf("%s:%s", y.Image, v)
	}
	return y.Command.GetImage()
}
