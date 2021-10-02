package docker

import (
	"satellite/internal/entity"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecGetContainerName(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetContainerName())

	exec.ContainerName = "container-name"
	assert.Equal(t, exec.GetContainerName(), "container-name")
}

func TestExecGetEnvFile(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetEnvFile())

	exec = Exec{EnvFile: "some_env_file"}
	assert.Equal(t, exec.GetEnvFile(), "--env-file some_env_file")
}

func TestExecGetUserId(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetUserId())

	exec = Exec{User: "501"}
	assert.Equal(t, exec.GetUserId(), "--user 501")
}

func TestExecGetWorkDir(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetWorkDir())

	exec = Exec{WorkDir: "some_work_dir"}
	assert.Equal(t, exec.GetWorkDir(), "--workdir=some_work_dir")
}

func TestExecGetDetached(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetDetached())

	exec = Exec{Detach: false}
	assert.Empty(t, exec.GetDetached())

	exec = Exec{Detach: true}
	assert.Equal(t, exec.GetDetached(), "-d")
}

func TestExecGetFlags(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetFlags())

	exec = Exec{Interactive: true}
	assert.Equal(t, exec.GetFlags(), "-i")

	exec = Exec{Tty: true}
	assert.Equal(t, exec.GetFlags(), "-t")

	exec.Interactive = true
	assert.Equal(t, exec.GetFlags(), "-it")

	exec.Detach = true
	assert.Empty(t, exec.GetFlags())
}

func TestExecGetEnv(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetEnv())

	exec = Exec{Env: []string{"MYVAR1=foo", "MYVAR2=bar"}}
	assert.Equal(t, exec.GetEnv(), "--env MYVAR1=foo --env MYVAR2=bar")
}

func TestExexGetBinBash(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetBinBash())

	exec = Exec{BinBash: true}
	assert.True(t, exec.GetBinBash())

	exec = Exec{BinBash: false}
	assert.False(t, exec.GetBinBash())
}

func TestExecGetPostCommands(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetPostCommands())

	exec = Exec{PostCommands: []string{"composer --version"}}
	assert.Equal(t, strings.Join(exec.GetPostCommands(), " "), "composer --version")

	exec.PostCommands = []string{"composer --version", "composer --version"}
	assert.Equal(t, strings.Join(exec.GetPostCommands(), " "), "composer --version; composer --version")
}

func TestExecGetPreCommands(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetPreCommands())

	exec = Exec{PreCommands: []string{"composer --version"}}
	assert.Equal(t, strings.Join(exec.GetPreCommands(), " "), "composer --version")

	exec.PreCommands = []string{"composer --version", "composer --version"}
	assert.Equal(t, strings.Join(exec.GetPreCommands(), " "), "composer --version; composer --version")
}

func TestExecToCommand(t *testing.T) {
	exec := Exec{}
	assert.Equal(t, strings.Join(exec.ToCommand([]string{}), " "), "exec")

	exec.Name = "some-name"
	exec.Interactive = true
	exec.Tty = true
	exec.ContainerName = "some-container-name"
	exec.WorkDir = "/some/work/dir"
	exec.PreCommands = []string{"pre command", "pre command 2"}
	result := strings.Join(exec.ToCommand([]string{"some", "command"}), " ")
	e := "exec -it --workdir=/some/work/dir some-container-name /bin/bash -c pre command; pre command 2; some command"
	assert.Equal(t, result, e)

	exec.Beginning = "php bin/console"
	exec.PreCommands = []string{"pre command", "pre command 2"}
	result = strings.Join(exec.ToCommand([]string{"cache", "clear"}), " ")
	e = "exec -it --workdir=/some/work/dir some-container-name /bin/bash -c pre command; pre command 2; php bin/console cache clear"
	assert.Equal(t, result, e)

	exec.PreCommands = nil
	result = strings.Join(exec.ToCommand([]string{"cache", "clear"}), " ")
	e = "exec -it --workdir=/some/work/dir some-container-name php bin/console cache clear"
	assert.Equal(t, result, e)
}

func TestExecGetExecCommand(t *testing.T) {
	exec := Exec{}
	assert.Equal(t, exec.GetExecCommand(), string(entity.DOCKER))
}

func TestExecGetName(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetName())

	exec.Name = "some-name"
	assert.Equal(t, exec.GetName(), "some-name")
}

func TestExecGetDescription(t *testing.T) {
	exec := Exec{}
	assert.Empty(t, exec.GetDescription())

	exec.Description = "some description"
	assert.Equal(t, exec.GetDescription(), "some description")
}
