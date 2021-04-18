package collector

import (
	"strings"
	"testing"

	"github.com/mamau/starter/libs"

	"github.com/mamau/starter/config"
	"github.com/mamau/starter/entity"
)

func TestCollectCommand(t *testing.T) {
	s := getService("service-with-pre-post-command", []string{"-v"})
	c := NewCollector(s)
	cc := c.CollectCommand()
	e := "run -ti service-with-pre-post-command:1 /bin/bash -c pre cmd 1; pre cmd 2; service -v; post cmd 1; post cmd 2"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s = getService("service-with-post-command", []string{"-v"})
	c = NewCollector(s)
	cc = c.CollectCommand()
	e = "run -ti service-with-post-command:1 /bin/bash -c service -v; post cmd 1; post cmd 2"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s = getService("service-with-pre-command", []string{"-v"})
	c = NewCollector(s)
	cc = c.CollectCommand()
	e = "run -ti service-with-pre-command:1 /bin/bash -c pre cmd 1; pre cmd 2; service -v"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s = getService("service", []string{"-v"})
	c = NewCollector(s)
	cc = c.CollectCommand()
	e = "run -ti service:1 service -v"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s.Config.Version = ""
	c = NewCollector(s)
	cc = c.CollectCommand()
	e = "run -ti service service -v"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s = getService("service-bin-bash", []string{"-v"})
	c = NewCollector(s)
	cc = c.CollectCommand()
	e = "run -ti service-bin-bash:1 /bin/bash -c service -v"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s = getService("full-service-docker-config", []string{"-v"})
	c = NewCollector(s)
	cc = c.CollectCommand()
	e = "run -ti -u 1000 -e SOME_VAR=someVal -e SOME_VAR2=someVal2 --add-host=docker:10.180.0.1 --add-host=somehost.com -p 80:80 -p 127.0.0.1.8089:8089 --dns=0.0.0.0 --dns=8.8.8.8 --workdir=/home/service -v /path/to/current/dir/cache:/tmp -v /path/to/current/dir:/home/service full-service-docker-config:2 /bin/bash -c pre cmd 1; pre cmd 2; service -v; post cmd 1; post cmd 2"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}
}

func TestClientCommand(t *testing.T) {
	s := getService("service-with-pre-post-command", []string{"-v"})
	c := NewCollector(s)
	cc := c.ClientCommand()
	e := "/bin/bash -c pre cmd 1; pre cmd 2; service -v; post cmd 1; post cmd 2"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong client command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s = getService("service-with-post-command", []string{"-v"})
	c = NewCollector(s)
	cc = c.ClientCommand()
	e = "/bin/bash -c service -v; post cmd 1; post cmd 2"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s = getService("service-with-pre-command", []string{"-v"})
	c = NewCollector(s)
	cc = c.ClientCommand()
	e = "/bin/bash -c pre cmd 1; pre cmd 2; service -v"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s = getService("service", []string{"-v"})
	c = NewCollector(s)
	cc = c.ClientCommand()
	e = "service -v"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s.Config.Version = ""
	c = NewCollector(s)
	cc = c.ClientCommand()
	e = "service -v"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}

	s = getService("service-bin-bash", []string{"-v"})
	c = NewCollector(s)
	cc = c.ClientCommand()
	e = "/bin/bash -c service -v"
	if strings.Join(cc, " ") != e {
		t.Errorf("wrong collect command, expected: %q\n got %q", e, strings.Join(cc, " "))
	}
}

func TestDockerConfigCommand(t *testing.T) {
	s := getService("service-docker-config", []string{"-v"})
	c := NewCollector(s)
	dcc := c.DockerConfigCommand()
	e := "-u 1000 -e SOME_VAR=someVal -e SOME_VAR2=someVal2 --add-host=docker:10.180.0.1 --add-host=somehost.com -p 80:80 -p 127.0.0.1.8089:8089 --dns=0.0.0.0 --dns=8.8.8.8 --workdir=/home/service -v /path/to/current/dir/cache:/tmp -v /path/to/current/dir:/home/service service-docker-config:2"
	if strings.Join(dcc, " ") != e {
		t.Errorf("wrong docker config command, expected: %q\n got %q", e, strings.Join(dcc, " "))
	}
}

func TestGetBeginCommand(t *testing.T) {
	s := getService("service-with-command", []string{"-v"})
	c := NewCollector(s)
	e := "exec -ti"
	if bc := c.getBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service", []string{"-v"})
	c = NewCollector(s)
	e = "run -ti"
	if bc := c.getBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-pull-cmd", []string{"-v"})
	c = NewCollector(s)
	e = "pull"
	if bc := c.getBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-flags", []string{"-v"})
	c = NewCollector(s)
	e = "run -T"
	if bc := c.getBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-flags-and-detached", []string{"-v"})
	c = NewCollector(s)
	e = "run -d"
	if bc := c.getBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-flags-and-pull", []string{"-v"})
	c = NewCollector(s)
	e = "pull"
	if bc := c.getBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-detach", []string{"-v"})
	c = NewCollector(s)
	e = "run -d"
	if bc := c.getBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-detach-and-pull", []string{"-v"})
	c = NewCollector(s)
	e = "pull"
	if bc := c.getBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-clean-up", []string{"-v"})
	c = NewCollector(s)
	e = "run -ti --rm"
	if bc := c.getBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}
}

func getService(n string, args []string) *entity.Service {
	c := setConfig()
	s := c.GetService(n)
	return entity.NewService(s, args)
}

func setConfig() *config.Config {
	config.NewConfig(libs.GetPwd() + "/testdata/starter")
	c := config.GetConfig()
	return c
}
