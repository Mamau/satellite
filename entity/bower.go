package entity

import (
	"sync"
)

var bOnce sync.Once
var bInstance *Bower

type Bower struct {
	*Command
}

//
//func NewBower(args []string) *Bower {
//	bOnce.Do(func() {
//		bInstance = &Bower{
//			Command: &Command{
//				CmdName:      "bower",
//				Image:        "mamau/bower",
//				HomeDir:      "/home/node",
//				Args:         args,
//				DockerConfig: config.GetConfig().GetBower(),
//			},
//		}
//	})
//
//	return bInstance
//}
//
//func (b *Bower) dockerCommandData() [][]string {
//	return [][]string{
//		b.DockerConfig.GetUserId(),
//		b.DockerConfig.GetEnvironmentVariables(),
//		b.DockerConfig.GetHosts(),
//		b.DockerConfig.GetPorts(),
//		b.DockerConfig.GetDns(),
//		b.workDir(),
//		b.cacheDir(),
//		b.projectVolume(),
//		{b.getImage()},
//	}
//}
//
//func (b *Bower) dockerDataToCommand() []string {
//	var fullCommand []string
//	for _, command := range b.dockerCommandData() {
//		fullCommand = append(fullCommand, command...)
//	}
//
//	return fullCommand
//}
//
//func (b *Bower) CollectCommand() []string {
//	return append(b.dockerDataToCommand(), b.fullCommand())
//}

func (b *Bower) getImage() string {
	return b.Image
}
