package entity

import (
	"github.com/mamau/starter/config/docker"
)

type Service struct {
	Config *docker.Docker
	*Command
}

func NewService(config *docker.Docker, args []string) *Service {
	homeDir := "/tmp"
	workDir := "/tmp"

	if config == nil {
		return nil
	}

	if wd := config.WorkDir; wd != "" {
		workDir = wd
	}
	if hd := config.HomeDir; hd != "" {
		homeDir = hd
	}

	return &Service{
		Config: config,
		Command: &Command{
			CmdName: config.Name,
			Image:   config.Name,
			HomeDir: homeDir,
			WorkDir: workDir,
			Version: config.Version,
			Args:    args,
		},
	}
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
