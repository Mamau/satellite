package entity

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/mamau/starter/libs"
)

var cOnce sync.Once
var cInstance *Composer

type Composer struct {
	*Command
}

func GetComposer(image, homeDir, version string, args []string) *Composer {
	cOnce.Do(func() {
		cInstance = &Composer{
			Command: &Command{
				Image:   image,
				HomeDir: homeDir,
				Version: version,
				Args:    args,
				Config:  libs.GetConfig().GetComposer(),
			},
		}
	})

	return cInstance
}

func (c *Composer) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		c.Config.GetDns(),
		c.workDirVolume(),
		c.projectVolume(),
		c.certsVolume(),
		c.cacheVolume(),
		{c.getImage()},
		{"/bin/bash", "-c", c.fullCommand()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}

func (c *Composer) fullCommand() string {
	configArgs := c.getConfigCommand()
	if configArgs != "" {
		configArgs += "; "
	}
	fullCommand := c.getConfigCommand() + c.getMainCommand() + c.getPostCommands()
	return fullCommand
}

func (c *Composer) getMainCommand() string {
	mainCommand := append([]string{"composer"}, c.Args...)
	return strings.Join(mainCommand, " ")
}

func (c *Composer) cacheVolume() []string {
	return []string{
		"-v",
		fmt.Sprintf("%s/cache/composer:/tmp", libs.GetPwd()),
	}
}

func (c *Composer) certsVolume() []string {
	if runtime.GOOS == "windows" {
		return []string{}
	}

	return []string{
		"-v",
		"/etc/ssl/certs:/etc/ssl/certs",
	}
}
