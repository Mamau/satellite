package collector

import (
	"strings"
	"testing"

	"github.com/mamau/starter/libs"

	"github.com/mamau/starter/config"
	"github.com/mamau/starter/entity"
)

func TestGetBeginCommand(t *testing.T) {
	s := getService("service-with-command", []string{"-v"})
	c := NewCollector(s)
	e := "exec -ti"
	if bc := c.GetBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service", []string{"-v"})
	c = NewCollector(s)
	e = "run -ti"
	if bc := c.GetBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-pull-cmd", []string{"-v"})
	c = NewCollector(s)
	e = "pull"
	if bc := c.GetBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-flags", []string{"-v"})
	c = NewCollector(s)
	e = "run -T"
	if bc := c.GetBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-flags-and-detached", []string{"-v"})
	c = NewCollector(s)
	e = "run -d"
	if bc := c.GetBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-flags-and-pull", []string{"-v"})
	c = NewCollector(s)
	e = "pull"
	if bc := c.GetBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-detach", []string{"-v"})
	c = NewCollector(s)
	e = "run -d"
	if bc := c.GetBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-detach-and-pull", []string{"-v"})
	c = NewCollector(s)
	e = "pull"
	if bc := c.GetBeginCommand(); strings.Join(bc, " ") != e {
		t.Errorf("begin command for service must be %q\n got %q", e, strings.Join(bc, " "))
	}

	s = getService("service-with-clean-up", []string{"-v"})
	c = NewCollector(s)
	e = "run -ti --rm"
	if bc := c.GetBeginCommand(); strings.Join(bc, " ") != e {
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
