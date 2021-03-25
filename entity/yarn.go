package entity

import (
	"sync"
)

var once sync.Once
var instance *Yarn

type Yarn struct {
	*Command
}

//func NewYarn(version string, args []string) *Yarn {
//	once.Do(func() {
//		instance = &Yarn{
//			Command: &Command{
//				CmdName:      "yarn",
//				Image:        "node",
//				HomeDir:      "/home/node",
//				Version:      version,
//				Args:         args,
//				DockerConfig: config.GetConfig().GetYarn(),
//			},
//		}
//	})
//
//	return instance
//}
//
//func (y *Yarn) CollectCommand() []string {
//	clientCmd := []string{"/bin/bash", "-c", y.fullCommand()}
//	return append(y.dockerDataToCommand(), clientCmd...)
//}
