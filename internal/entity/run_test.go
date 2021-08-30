package entity

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCommand(t *testing.T) {
	run := Run{}
	result := strings.Join(run.ToCommand([]string{}), " ")
	assert.Equal(t, result, "run")

	run.Detach = true
	run.ContainerName = "some-name"
	run.Image = "my-img"
	run.Version = "12"
	run.PreCommands = []string{"command --version", "command -v"}
	run.PostCommands = []string{"exit"}
	run.Dns = []string{"0.0.0.0"}

	result = strings.Join(run.ToCommand([]string{"pwd"}), " ")
	assert.Equal(t, result, "run -d --dns=0.0.0.0 --name some-name my-img:12 /bin/bash -c command --version; command -v; pwd; exit")

	run = Run{
		Image:   "test",
		Version: "1",
		Dns:     []string{"0.0.8.8"},
	}
	result = strings.Join(run.ToCommand([]string{"ls"}), " ")
	assert.Equal(t, result, "run --dns=0.0.8.8 test:1 ls")

	run = Run{
		Image:     "composer",
		Version:   "1.9",
		Beginning: "composer",
	}
	result = strings.Join(run.ToCommand([]string{"install", "--ignore-platform-reqs"}), " ")
	assert.Equal(t, result, "run composer:1.9 composer install --ignore-platform-reqs")
}

func TestGetExecCommand(t *testing.T) {
	run := Run{}
	assert.Equal(t, run.GetExecCommand(), string(DOCKER))
}

func TestGetBinBash(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetBinBash())

	run = Run{BinBash: true}
	assert.True(t, run.GetBinBash())

	run = Run{BinBash: false}
	assert.False(t, run.GetBinBash())
}

func TestGetPostCommands(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetPostCommands())

	run = Run{PostCommands: []string{"composer --version"}}
	assert.Equal(t, strings.Join(run.GetPostCommands(), " "), "composer --version")

	run.PostCommands = []string{"composer --version", "composer --version"}
	assert.Equal(t, strings.Join(run.GetPostCommands(), " "), "composer --version; composer --version")
}

func TestGetPreCommands(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetPreCommands())

	run = Run{PreCommands: []string{"composer --version"}}
	assert.Equal(t, strings.Join(run.GetPreCommands(), " "), "composer --version")

	run.PreCommands = []string{"composer --version", "composer --version"}
	assert.Equal(t, strings.Join(run.GetPreCommands(), " "), "composer --version; composer --version")
}

func TestGetImage(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetImage())

	run = Run{Image: "composer"}
	assert.Equal(t, run.GetImage(), "composer")

	run.Version = "2"
	assert.Equal(t, run.GetImage(), "composer:2")
}

func TestGetContainerName(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetContainerName())

	run = Run{Name: "some-name"}
	assert.Equal(t, run.GetContainerName(), "--name some-name")

	run.ContainerName = "container-name"
	assert.Equal(t, run.GetContainerName(), "--name container-name")
}

func TestGetName(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetName())

	run = Run{Name: "some-name"}
	assert.Equal(t, run.GetName(), "some-name")
}

func TestGetVolumes(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetVolumes())

	run = Run{Volumes: []string{"/path/project:/var/www", "/path/project/log:/var/log"}}
	assert.Equal(t, run.GetVolumes(), "-v /path/project:/var/www -v /path/project/log:/var/log")
}

func TestGetDns(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetDns())

	run = Run{Dns: []string{"0.0.0.0", "8.8.8.8"}}
	assert.Equal(t, run.GetDns(), "--dns=0.0.0.0 --dns=8.8.8.8")
}

func TestGetPorts(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetPorts())

	run = Run{Ports: []string{"127.0.0.1:80:8080/tcp", "443:433"}}
	assert.Equal(t, run.GetPorts(), "-p 127.0.0.1:80:8080/tcp -p 443:433")
}

func TestGetHosts(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetHosts())

	run = Run{AddHosts: []string{"docker:10.180.0.1", "anyHost"}}
	assert.Equal(t, run.GetHosts(), "--add-host=docker:10.180.0.1 --add-host=anyHost")
}

func TestGetEnv(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetEnv())

	run = Run{Env: []string{"MYVAR1=foo", "MYVAR2=bar"}}
	assert.Equal(t, run.GetEnv(), "--env MYVAR1=foo --env MYVAR2=bar")
}

func TestGetUserId(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetUserId())

	run = Run{User: "501"}
	assert.Equal(t, run.GetUserId(), "--user 501")
}

func TestGetNetwork(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetNetwork())

	run = Run{Network: "my_network"}
	assert.Equal(t, run.GetNetwork(), "--network my_network")
}

func TestGetCleanUp(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetCleanUp())

	run = Run{CleanUp: true}
	assert.Equal(t, run.GetCleanUp(), "--rm")
}

func TestGetFlags(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetFlags())

	run = Run{Interactive: true}
	assert.Equal(t, run.GetFlags(), "-i")

	run = Run{Tty: true}
	assert.Equal(t, run.GetFlags(), "-t")

	run.Interactive = true
	assert.Equal(t, run.GetFlags(), "-it")

	run.Detach = true
	assert.Empty(t, run.GetFlags())
}

func TestGetWorkDir(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetWorkDir())

	run = Run{WorkDir: "some_work_dir"}
	assert.Equal(t, run.GetWorkDir(), "--workdir=some_work_dir")
}

func TestGetHostname(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetHostname())

	run = Run{Hostname: "some_host_name"}
	assert.Equal(t, run.GetHostname(), "--hostname some_host_name")
}

func TestGetEnvFile(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetEnvFile())

	run = Run{EnvFile: "some_env_file"}
	assert.Equal(t, run.GetEnvFile(), "--env-file some_env_file")
}

func TestGetDetached(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetDetached())

	run = Run{Detach: false}
	assert.Empty(t, run.GetDetached())

	run = Run{Detach: true}
	assert.Equal(t, run.GetDetached(), "-d")
}

func TestGetDescription(t *testing.T) {
	run := Run{}
	assert.Empty(t, run.GetDescription())

	run = Run{Description: "install composer dependencies"}
	assert.Equal(t, run.GetDescription(), "install composer dependencies")
}
