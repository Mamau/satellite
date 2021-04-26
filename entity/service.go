package entity

import (
	"fmt"
	"strings"

	"github.com/mamau/satellite/config/docker"
)

type Service struct {
	Config *docker.Docker
	Args   []string
}

func NewService(config *docker.Docker, args []string) *Service {
	//TODO: need validation
	return &Service{
		Config: config,
		Args:   args,
	}
}

func (s *Service) GetWorkDir() string {
	if s.Config.WorkDir != "" {
		return fmt.Sprintf("--workdir=%s", s.Config.WorkDir)
	}

	if s.Config.HomeDir != "" {
		return fmt.Sprintf("--workdir=%s", s.Config.HomeDir)
	}

	return ""
}

func (s *Service) GetClientCommand() string {
	if s.Config.ImageCommand == "" {
		return ""
	}

	if len(s.Config.GetPreCommands()) > 0 || len(s.Config.GetPostCommands()) > 0 {
		cmd := append([]string{s.Config.ImageCommand}, s.Args...)
		return strings.Join(cmd, " ")
	}

	return strings.Join(s.Args, " ")
}

func (s *Service) GetImage() string {
	imageName := s.Config.Image
	if imageName == "" {
		imageName = s.Config.Name
	}

	if s.Config.Version != "" {
		return fmt.Sprintf("%s:%s", imageName, s.Config.Version)
	}
	return imageName
}

func (s *Service) GetImageCommand() string {
	if s.Config.ImageCommand == "" {
		return ""
	}

	if len(s.Config.GetPreCommands()) > 0 || len(s.Config.GetPostCommands()) > 0 {
		return "/bin/bash -c"
	}

	if s.Config.BinBash == true {
		return "/bin/bash -c " + s.Config.ImageCommand
	}

	return s.Config.ImageCommand
}

func (s *Service) GetDockerConfig() *docker.Docker {
	return s.Config
}
