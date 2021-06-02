package entity

import (
	"fmt"
	"strings"

	docker2 "github.com/mamau/satellite/internal/config/docker"
)

type Service struct {
	Config *docker2.Docker
	Args   []string
}

func NewService(config *docker2.Docker, args []string) *Service {
	return &Service{
		Config: config,
		Args:   args,
	}
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

func (s *Service) GetDockerConfig() *docker2.Docker {
	return s.Config
}
