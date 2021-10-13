package docker_compose

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpToCommand(t *testing.T) {
	up := Up{}
	up.Path = "./some/path"
	up.Detach = true
	up.NoDeps = true
	up.NoBuild = true

	e := "--file ./some/path up -d --no-deps --no-build"
	assert.Equal(t, e, strings.Join(up.ToCommand([]string{}), " "))
}

func TestGetRunNoDeps(t *testing.T) {
	up := Up{}
	assert.Empty(t, up.GetNoDeps())

	up.NoDeps = true
	assert.Equal(t, "--no-deps", up.GetNoDeps())
}

func TestGetUpDetached(t *testing.T) {
	up := Up{}
	assert.Empty(t, up.GetDetached())

	up.Detach = true
	assert.Equal(t, "-d", up.GetDetached())
}

func TestGetNoStart(t *testing.T) {
	up := Up{}
	assert.Empty(t, up.GetNoStart())

	up.NoStart = true
	assert.Equal(t, "--no-start", up.GetNoStart())
}

func TestGetUpRemoveOrphans(t *testing.T) {
	up := Up{}
	assert.Empty(t, up.GetRemoveOrphans())

	up.RemoveOrphans = true
	assert.Equal(t, "--remove-orphans", up.GetRemoveOrphans())
}

func TestGetNoBuild(t *testing.T) {
	up := Up{}
	assert.Empty(t, up.GetNoBuild())

	up.NoBuild = true
	assert.Equal(t, "--no-build", up.GetNoBuild())
}

func TestGetBuild(t *testing.T) {
	up := Up{}
	assert.Empty(t, up.GetBuild())

	up.Build = true
	assert.Equal(t, "--build", up.GetBuild())
}
