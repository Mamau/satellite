package entity

import (
	"github.com/mamau/starter/config/docker"
)

type Service struct {
	Config *docker.Docker
	*Command
}

func NewService(config *docker.Docker, args []string) *Service {
	if config == nil {
		return nil
	}

	return &Service{
		Config: config,
		Command: &Command{
			CmdName: config.GetClientCommand(),
			Image:   config.Name,
			Version: config.Version,
			Args:    args,
		},
	}
}

func (s *Service) GetProjectVolume() string {
	return ""
}

func (s *Service) GetDockerConfig() *docker.Docker {
	return s.Config
}

func (s *Service) GetCommandConfig() *Command {
	return s.Command
}

func (s *Service) GetClientSignature(cmd []string) []string {
	return cmd
}
