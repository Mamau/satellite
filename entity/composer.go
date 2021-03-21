package entity

import (
	"strings"
	"sync"

	"github.com/mamau/starter/libs/services/composer"

	"github.com/mamau/starter/libs"
)

var cOnce sync.Once
var cInstance *Composer

type Composer struct {
	*composer.Config
	*Command
}

func NewComposer(version string, args []string) *Composer {
	cOnce.Do(func() {
		cInstance = &Composer{
			Config: &libs.GetConfig().GetComposer().Config,
			Command: &Command{
				CmdName:      "composer",
				Image:        "composer",
				HomeDir:      "/home/www-data",
				Version:      version,
				Args:         args,
				DockerConfig: libs.GetConfig().GetComposer(),
			},
		}
	})

	return cInstance
}

func (c *Composer) CollectCommand() []string {
	clientCmd := []string{"/bin/bash", "-c", c.fullCommand()}
	return append(c.dockerDataToCommand(), clientCmd...)
}

func (c *Composer) fullCommand() string {
	return c.configToCommand() + c.getPreCommands() + c.getCommand() + c.getPostCommands()
}

func (c *Composer) configToCommand() string {
	configCommands := libs.DeleteEmpty(c.Config.GetAll())
	if all := configCommands; all != nil {
		return strings.Join(configCommands, "; ") + "; "
	}

	return ""
}
