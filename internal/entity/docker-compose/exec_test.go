package docker_compose

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecToCommand(t *testing.T) {
	exec := Exec{}
	exec.Path = "./some/path"
	exec.Detach = true
	exec.Env = []string{
		"VAR1=VAL1",
		"VAR2=VAL2",
	}

	e := "--file ./some/path exec -d -e VAR1=VAL1 -e VAR2=VAL2"
	assert.Equal(t, e, strings.Join(exec.ToCommand([]string{}), " "))
}

func TestGetDetached(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetDetached())

	exec.Detach = true
	assert.Equal(t, "-d", exec.GetDetached())
}

func TestGetUserId(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetUserId())

	exec.User = "1000"
	assert.Equal(t, "-u 1000", exec.GetUserId())
}

func TestGetEnv(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetEnv())

	exec.Env = []string{
		"VAR1=VAL1",
		"VAR2=VAL2",
	}
	assert.Equal(t, "-e VAR1=VAL1 -e VAR2=VAL2", exec.GetEnv())
}

func TestGetWorkdir(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetWorkdir())

	exec.Workdir = "/some/dir"
	assert.Equal(t, "--workdir /some/dir", exec.GetWorkdir())
}
