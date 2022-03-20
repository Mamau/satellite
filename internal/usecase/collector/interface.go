package collector

import "satellite/internal/entity"

//nolint
//go:generate mockgen -destination=../../mock/finder_mock.go -package=mock satellite/internal/usecase/collector Finder
type Finder interface {
	FindCommand(name string) entity.Runner
}
