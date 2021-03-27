package entity

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/mamau/starter/config/docker"

	"github.com/mamau/starter/config"
	"github.com/mamau/starter/config/composer"

	"github.com/mamau/starter/libs"
)

var cOnce sync.Once
var cInstance *Composer

type Composer struct {
	Config *composer.Composer
	*Command
}

func NewComposer(version string, args []string) *Composer {
	cOnce.Do(func() {
		cInstance = &Composer{
			Config: config.GetConfig().GetComposer(),
			Command: &Command{
				CmdName: "composer",
				Image:   "composer",
				HomeDir: "/home/www-data",
				Version: version,
				Args:    args,
			},
		}
	})

	return cInstance
}

func (c *Composer) GetDockerConfig() *docker.Docker {
	if c.Config == nil {
		return nil
	}
	if reflect.DeepEqual(c.Config.Docker, docker.Docker{}) {
		return nil
	}
	return &c.Config.Docker
}

func (c *Composer) GetCommandConfig() *Command {
	return c.Command
}

func (c *Composer) GetClientSignature(cmd []string) []string {
	command := append(c.configToCommand(), cmd...)
	return []string{"/bin/bash", "-c", strings.Join(command, " ")}
}

func (c *Composer) configToCommand() []string {
	if c.Config.Config == nil {
		return []string{}
	}
	configCommands := libs.DeleteEmpty(c.Config.GetAll())
	if len(configCommands) == 0 {
		return []string{}
	}

	for i, v := range configCommands {
		configCommands[i] = v + ";"
	}

	return libs.MergeSliceOfString(configCommands)
}

func (c *Composer) GetImage() string {
	if c.Config.Config == nil {
		return ""
	}
	if v := c.Config.GetVersion(); v != "" {
		return fmt.Sprintf("%s:%s", c.Image, v)
	}
	return c.Command.GetImage()
}
