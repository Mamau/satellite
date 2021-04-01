package entity

import (
	"fmt"
	"strings"

	"github.com/mamau/starter/libs"

	"github.com/mamau/starter/config/yarn"

	"github.com/mamau/starter/config/docker"

	"github.com/mamau/starter/config"
)

type Yarn struct {
	Config *yarn.Yarn
	*Command
}

func NewYarn(version string, args []string) *Yarn {
	return &Yarn{
		Config: config.GetConfig().GetYarn(),
		Command: &Command{
			CmdName: "yarn",
			Image:   "node",
			HomeDir: "/home/node",
			Version: version,
			Args:    args,
		},
	}
}

func (y *Yarn) GetDockerConfig() *docker.Docker {
	if y.Config == nil {
		return nil
	}
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
	if y.Config == nil {
		return []string{}
	}

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
	if y.Config == nil {
		return y.Command.GetImage()
	}
	if v := y.Config.GetVersion(); v != "" {
		return fmt.Sprintf("%s:%s", y.Image, v)
	}
	return y.Command.GetImage()
}
