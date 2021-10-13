package docker_compose

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildToCommand(t *testing.T) {
	build := Build{}
	build.Path = "./some/path"
	build.BuildArgs = []string{
		"KEY1=VAL1",
		"KEY2=VAL2",
	}
	build.Compress = true
	build.NoRm = true
	build.NoCache = true
	e := "--file ./some/path build --build-arg KEY1=VAL1 --build-arg KEY2=VAL2 --compress --no-cache --no-rm"
	assert.Equal(t, e, strings.Join(build.ToCommand([]string{}), " "))
}

func TestGetBuildArgs(t *testing.T) {
	build := Build{}
	assert.Empty(t, build.GetBuildArgs())

	build.BuildArgs = []string{
		"KEY1=VAL1",
		"KEY2=VAL2",
	}
	assert.Equal(t, "--build-arg KEY1=VAL1 --build-arg KEY2=VAL2", build.GetBuildArgs())
}

func TestGetCompress(t *testing.T) {
	build := Build{}
	assert.Empty(t, build.GetCompress())

	build.Compress = true
	assert.Equal(t, "--compress", build.GetCompress())
}

func TestGetForceRm(t *testing.T) {
	build := Build{}
	assert.Empty(t, build.GetForceRm())

	build.ForceRm = true
	assert.Equal(t, "--force-rm", build.GetForceRm())
}

func TestGetMemory(t *testing.T) {
	build := Build{}
	assert.Empty(t, build.GetMemory())

	build.Memory = "1000"
	assert.Equal(t, "--memory 1000", build.GetMemory())
}

func TestGetNoCache(t *testing.T) {
	build := Build{}
	assert.Empty(t, build.GetNoCache())

	build.NoCache = true
	assert.Equal(t, "--no-cache", build.GetNoCache())
}

func TestGetNoRm(t *testing.T) {
	build := Build{}
	assert.Empty(t, build.GetNoRm())

	build.NoRm = true
	assert.Equal(t, "--no-rm", build.GetNoRm())
}

func TestGetParallel(t *testing.T) {
	build := Build{}
	assert.Empty(t, build.GetParallel())

	build.Parallel = true
	assert.Equal(t, "--parallel", build.GetParallel())
}

func TestGetQuiet(t *testing.T) {
	build := Build{}
	assert.Empty(t, build.GetQuiet())

	build.Quiet = true
	assert.Equal(t, "--quiet", build.GetQuiet())
}

func TestGetPull(t *testing.T) {
	build := Build{}
	assert.Empty(t, build.GetPull())

	build.Pull = true
	assert.Equal(t, "--pull", build.GetPull())
}
