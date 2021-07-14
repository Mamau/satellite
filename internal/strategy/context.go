package strategy

import (
	"context"

	"github.com/mamau/satellite/internal/config/docker"
)

const KEY = "config"

type CommandContext struct {
	context.Context
}

func ContextWithConfig(parent context.Context, config *docker.Docker) CommandContext {
	ctx := context.WithValue(parent, KEY, *config)
	return CommandContext{
		ctx,
	}
}

func (c CommandContext) GetConfig() *docker.Docker {
	config := c.Value(KEY).(docker.Docker)

	return &config
}
