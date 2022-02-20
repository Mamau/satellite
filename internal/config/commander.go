package config

import "satellite/internal/entity"

type Commander interface {
	GetCommands() []entity.Runner
}
