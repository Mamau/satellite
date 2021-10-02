package docker_compose

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunToCommand(t *testing.T) {
	run := Run{}
	run.Path = "./some/path"
	run.Detach = true
	run.Env = []string{
		"VAR1=VAL1",
		"VAR2=VAL2",
	}

	e := "--file ./some/path run -d -e VAR1=VAL1 -e VAR2=VAL2 any val"
	assert.Equal(t, e, strings.Join(run.ToCommand([]string{"any", "val"}), " "))
}

func TestRunGetDetached(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetDetached())

	run.Detach = true
	assert.Equal(t, "-d", run.GetDetached())
}

func TestGetNoDeps(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetNoDeps())

	run.NoDeps = true
	assert.Equal(t, "--no-deps", run.GetNoDeps())
}

func TestRunGetUserId(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetUserId())

	run.User = "1000"
	assert.Equal(t, "-u 1000", run.GetUserId())
}

func TestGetEntryPoint(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetEntryPoint())

	run.Entrypoint = "/some/point"
	assert.Equal(t, "--entrypoint /some/point", run.GetEntryPoint())
}

func TestRunGetEnv(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetEnv())

	run.Env = []string{
		"VAR1=VAL1",
		"VAR2=VAL2",
	}
	assert.Equal(t, "-e VAR1=VAL1 -e VAR2=VAL2", run.GetEnv())
}

func TestGetPorts(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetPorts())

	run.Ports = []string{
		"80:80",
		"9001:9000",
	}
	assert.Equal(t, "-p 80:80 -p 9001:9000", run.GetPorts())
}
