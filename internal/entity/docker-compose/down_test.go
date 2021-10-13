package docker_compose

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownToCommand(t *testing.T) {
	down := Down{}
	down.Path = "./some/path"
	down.RemoveOrphans = true
	down.Rmi = "all"
	down.RemoveVolumes = true

	e := "--file ./some/path down --remove-orphans --rmi all --volumes"
	assert.Equal(t, e, strings.Join(down.ToCommand([]string{}), " "))
}

func TestGetRemoveOrphans(t *testing.T) {
	down := Down{}
	assert.Empty(t, down.GetRemoveOrphans())

	down.RemoveOrphans = true
	assert.Equal(t, "--remove-orphans", down.GetRemoveOrphans())
}

func TestGetRmi(t *testing.T) {
	down := Down{}
	assert.Empty(t, down.GetRmi())

	down.Rmi = "local"
	assert.Equal(t, "--rmi local", down.GetRmi())
}

func TestGetRemoveVolumes(t *testing.T) {
	down := Down{}
	assert.Empty(t, down.GetRemoveVolumes())

	down.RemoveVolumes = true
	assert.Equal(t, "--volumes", down.GetRemoveVolumes())
}
