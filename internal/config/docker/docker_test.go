package docker

import (
	"fmt"
	"os/user"
	"strings"
	"testing"
)

func TestGetImageCommand(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "image-command", docker.GetImageCommand)

	docker.ImageCommand = "--version"
	e := "--version"
	r := docker.GetImageCommand()
	if e != r {
		t.Errorf("expected image-command is %q, got %q", e, r)
	}

	docker.BinBash = true
	e = "/bin/bash -c --version"
	r = docker.GetImageCommand()
	if e != r {
		t.Errorf("expected image-command is %q, got %q", e, r)
	}

	docker.BinBash = false
	docker.PreCommands = []string{"ls -la"}
	e = "/bin/bash -c"
	r = docker.GetImageCommand()
	if e != r {
		t.Errorf("when bin-bash is false expected image-command is %q, got %q", e, r)
	}

	docker.PreCommands = nil
	docker.PostCommands = []string{"ls -la"}
	e = "/bin/bash -c"
	r = docker.GetImageCommand()
	if e != r {
		t.Errorf("when bin-bash is false expected image-command is %q, got %q", e, r)
	}

	docker.BinBash = true
	docker.PostCommands = []string{"ls -la"}
	e = "/bin/bash -c"
	r = docker.GetImageCommand()
	if e != r {
		t.Errorf("when bin-bash is true expected image-command is %q, got %q", e, r)
	}

	docker.PreCommands = nil
	docker.PostCommands = []string{"ls -la"}
	e = "/bin/bash -c"
	r = docker.GetImageCommand()
	if e != r {
		t.Errorf("when bin-bash is true expected image-command is %q, got %q", e, r)
	}
}

func TestGetImage(t *testing.T) {
	docker := Docker{}
	docker.Name = "imgName"
	e := "imgName"
	r := docker.GetImage()
	if e != r {
		t.Errorf("expected image is %q, got %q", e, r)
	}
	docker.Image = "node"
	e = "node"
	r = docker.GetImage()
	if e != r {
		t.Errorf("expected image is %q, got %q", e, r)
	}
	docker.Version = "10"
	e = "node:10"
	r = docker.GetImage()
	if e != r {
		t.Errorf("expected image is %q, got %q", e, r)
	}
}

func TestGetDns(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "dns", docker.GetDns)

	docker.Dns = []string{"8.8.8.8"}
	e := "--dns=8.8.8.8"
	r := docker.GetDns()
	if e != r {
		t.Errorf("expected dns is %q, got %q", e, r)
	}

	docker.Dns = []string{"8.8.8.8", "8.8.4.4"}
	e = "--dns=8.8.8.8 --dns=8.8.4.4"
	r = docker.GetDns()
	if e != r {
		t.Errorf("expected dns is %q, got %q", e, r)
	}
}

func TestGetVolumes(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "volumes", docker.GetVolumes)

	docker.Volumes = []string{"/tmp:/var/www"}
	e := "-v /tmp:/var/www"
	r := docker.GetVolumes()
	if e != r {
		t.Errorf("expected volumes is %q, got %q", e, r)
	}

	docker.Volumes = []string{"/tmp:/var/www", "/var/log:/var/www"}
	e = "-v /tmp:/var/www -v /var/log:/var/www"
	r = docker.GetVolumes()
	if e != r {
		t.Errorf("expected volumes is %q, got %q", e, r)
	}
}

func TestGetPorts(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "ports", docker.GetPorts)

	docker.Ports = []string{"80:80"}
	e := "-p 80:80"
	r := docker.GetPorts()
	if e != r {
		t.Errorf("expected ports is %q, got %q", e, r)
	}

	docker.Ports = []string{"80:80", "127.0.0.1:443:443"}
	e = "-p 80:80 -p 127.0.0.1:443:443"
	r = docker.GetPorts()
	if e != r {
		t.Errorf("expected ports is %q, got %q", e, r)
	}
}

func TestGetHosts(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "add-hosts", docker.GetHosts)

	docker.AddHosts = []string{"somehost.com"}
	e := "--add-host=somehost.com"
	r := docker.GetHosts()
	if e != r {
		t.Errorf("expected add-hosts is %q, got %q", e, r)
	}

	docker.AddHosts = []string{"somehost.com", "127.0.0.1"}
	e = "--add-host=somehost.com --add-host=127.0.0.1"
	r = docker.GetHosts()
	if e != r {
		t.Errorf("expected add-hosts is %q, got %q", e, r)
	}
}

func TestGetEnvironmentVariables(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "environment-variables", docker.GetEnvironmentVariables)

	docker.EnvVars = []string{"PHP_IDE_CONFIG=serverName=192.168.0.1"}
	e := "-e PHP_IDE_CONFIG=serverName=192.168.0.1"
	r := docker.GetEnvironmentVariables()
	if e != r {
		t.Errorf("expected environment-variables is %q, got %q", e, r)
	}

	docker.EnvVars = []string{"PHP_IDE_CONFIG=serverName=192.168.0.1", "SOME_VAR=SOME_VAL"}
	e = "-e PHP_IDE_CONFIG=serverName=192.168.0.1 -e SOME_VAR=SOME_VAL"
	r = docker.GetEnvironmentVariables()
	if e != r {
		t.Errorf("expected environment-variables is %q, got %q", e, r)
	}
}

func TestGetUserId(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "user-id", docker.GetUserId)

	docker.UserId = "$(uid)"
	currentUser, err := user.Current()
	if err != nil {
		t.Errorf("error while getting current user, err: %q", err)
	}
	e := fmt.Sprintf("-u %s", currentUser.Uid)
	r := docker.GetUserId()
	if e != r {
		t.Errorf("expected user-id is %q, got %q", e, r)
	}

	docker.UserId = "1000"
	e = "-u 1000"
	r = docker.GetUserId()
	if e != r {
		t.Errorf("expected user-id is %q, got %q", e, r)
	}
}

func TestGetWorkDir(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "work-dir", docker.GetWorkDir)

	docker.HomeDir = "/home/dir"
	e := "--workdir=/home/dir"
	r := docker.GetWorkDir()
	if e != r {
		t.Errorf("expected work-dir is %q, got %q", e, r)
	}

	docker.WorkDir = "/work/dir"
	e = "--workdir=/work/dir"
	r = docker.GetWorkDir()
	if e != r {
		t.Errorf("expected work-dir is %q, got %q", e, r)
	}
}

func TestGetPostCommands(t *testing.T) {
	docker := Docker{}
	r := strings.Join(docker.GetPostCommands(), " ")
	e := ""
	if e != r {
		t.Errorf("expected post-commands is %q, got %q", e, r)
	}

	docker.PostCommands = []string{"some command"}
	e = "some command"
	r = strings.Join(docker.GetPostCommands(), " ")
	if e != r {
		t.Errorf("expected post-commands is %q, got %q", e, r)
	}

	docker.PostCommands = []string{"some command", "some command2"}
	e = "some command; some command2"
	r = strings.Join(docker.GetPostCommands(), " ")
	if e != r {
		t.Errorf("expected post-commands is %q, got %q", e, r)
	}
}

func TestGetPreCommands(t *testing.T) {
	docker := Docker{}
	r := strings.Join(docker.GetPreCommands(), " ")
	e := ""
	if e != r {
		t.Errorf("expected pre-commands is %q, got %q", e, r)
	}

	docker.PreCommands = []string{"some command"}
	e = "some command"
	r = strings.Join(docker.GetPreCommands(), " ")
	if e != r {
		t.Errorf("expected pre-commands is %q, got %q", e, r)
	}

	docker.PreCommands = []string{"some command", "some command2"}
	e = "some command; some command2"
	r = strings.Join(docker.GetPreCommands(), " ")
	if e != r {
		t.Errorf("expected pre-commands is %q, got %q", e, r)
	}
}

func TestGetFlags(t *testing.T) {
	docker := Docker{}
	e := "-ti"
	r := docker.GetFlags()
	if e != r {
		t.Errorf("expected flags is %q, got %q", e, r)
	}
	docker.Flags = "-T"
	e = "-T"
	r = docker.GetFlags()
	if e != r {
		t.Errorf("expected flags is %q, got %q", e, r)
	}
}

func TestGetDetached(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "detach", docker.GetDetached)

	docker.Detach = true
	e := "-d"
	r := docker.GetDetached()
	if e != r {
		t.Errorf("expected detach is %q, got %q", e, r)
	}

	docker.Detach = false
	checkForEmpty(t, "detach", docker.GetDetached)
}

func TestGetDockerCommand(t *testing.T) {
	docker := Docker{}
	e := "run"
	r := docker.GetDockerCommand()
	if e != r {
		t.Errorf("expected command is %q, got %q", e, r)
	}

	docker.Command = "pull"
	e = "pull"
	r = docker.GetDockerCommand()
	if e != r {
		t.Errorf("expected command is %q, got %q", e, r)
	}
}

func TestGetCleanUp(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "clean-up", docker.GetCleanUp)

	docker.CleanUp = true
	e := "--rm"
	r := docker.GetCleanUp()
	if e != r {
		t.Errorf("expected clean-up is %q, got %q", e, r)
	}

	docker.CleanUp = false
	checkForEmpty(t, "clean-up", docker.GetCleanUp)
}

func TestGetContainerName(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "container name", docker.GetContainerName)

	docker.ContainerName = "test_container"
	e := "--name test_container"
	r := docker.GetContainerName()
	if e != r {
		t.Errorf("expected container name is %q, got %q", e, r)
	}
}

func TestGetNetwork(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "network", docker.GetNetwork)

	docker.Network = "test-network"
	e := "--network test-network"
	r := docker.GetNetwork()
	if e != r {
		t.Errorf("expected network %q, got %q", e, r)
	}
}

func checkForEmpty(t *testing.T, p string, c func() string) {
	e := ""
	r := c()
	if e != r {
		t.Errorf("expected %q is empty %q, got %q", p, e, r)
	}
}
