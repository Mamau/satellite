package docker

import (
	"fmt"
	"os/user"
	"testing"
)

//func TestGetEnvironmentVariables(t *testing.T) {
//	docker := Docker{}
//	checkForEmpty(t, "environment-variables", docker.GetEnvironmentVariables)
//
//	docker.EnvVars = []string{"PHP_IDE_CONFIG=serverName=192.168.0.1"}
//
//}

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
	checkForEmpty(t, "post-commands", docker.GetPostCommands)

	docker.PostCommands = []string{"some command"}
	e := "some command"
	r := docker.GetPostCommands()
	if e != r {
		t.Errorf("expected post-commands is %q, got %q", e, r)
	}

	docker.PostCommands = []string{"some command", "some command2"}
	e = "some command; some command2"
	r = docker.GetPostCommands()
	if e != r {
		t.Errorf("expected post-commands is %q, got %q", e, r)
	}
}

func TestGetPreCommands(t *testing.T) {
	docker := Docker{}
	checkForEmpty(t, "pre-commands", docker.GetPreCommands)

	docker.PreCommands = []string{"some command"}
	e := "some command"
	r := docker.GetPreCommands()
	if e != r {
		t.Errorf("expected pre-commands is %q, got %q", e, r)
	}

	docker.PreCommands = []string{"some command", "some command2"}
	e = "some command; some command2"
	r = docker.GetPreCommands()
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
