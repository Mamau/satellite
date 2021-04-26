package entity

import (
	"testing"

	"github.com/mamau/satellite/config"
	"github.com/mamau/satellite/libs"
)

func TestGetImage(t *testing.T) {
	s := getService("service-with-image", []string{"-v"})
	img := s.GetImage()
	e := "some-service-image:7.4"
	if img != e {
		t.Errorf("image must be %q, got %q", e, img)
	}

	s.GetDockerConfig().Version = ""
	img = s.GetImage()
	e = "some-service-image"
	if img != e {
		t.Errorf("image must be %q, got %q", e, img)
	}

	s.GetDockerConfig().Image = ""
	s.GetDockerConfig().Version = "7.1"
	img = s.GetImage()
	e = "service-with-image:7.1"
	if img != e {
		t.Errorf("image must be %q, got %q", e, img)
	}

	s.GetDockerConfig().Image = ""
	s.GetDockerConfig().Version = ""
	img = s.GetImage()
	e = "service-with-image"
	if img != e {
		t.Errorf("image must be %q, got %q", e, img)
	}
}

func TestGetClientCommand(t *testing.T) {
	s := getService("service", []string{"-v"})
	cc := s.GetClientCommand()
	e := "-v"
	if cc != e {
		t.Errorf("service command must be %q, got %q", e, cc)
	}

	s.GetDockerConfig().PreCommands = []string{"pre", "command"}
	cc = s.GetClientCommand()
	e = "service -v"
	if cc != e {
		t.Errorf("service command must be %q, got %q", e, cc)
	}

	s.GetDockerConfig().PreCommands = nil
	s.GetDockerConfig().PostCommands = []string{"post", "command"}
	cc = s.GetClientCommand()
	e = "service -v"
	if cc != e {
		t.Errorf("service command must be %q, got %q", e, cc)
	}

	s.GetDockerConfig().ImageCommand = ""
	cc = s.GetClientCommand()
	e = ""
	if cc != e {
		t.Errorf("service command must be empty, got %q", cc)
	}
}

func TestGetFlags(t *testing.T) {
	s := getService("service-flags", []string{"-v"})
	if c := s.GetDockerConfig().GetFlags(); c != "-T" {
		t.Errorf("service must have glag %q,\n got %q\n", "-T", c)
	}
	s.GetDockerConfig().Command = "pull"
	if c := s.GetDockerConfig().GetFlags(); c != "" {
		t.Errorf("service must be empty, got %q\n", c)
	}
	s.GetDockerConfig().Command = ""
	s.GetDockerConfig().Detach = true
	if c := s.GetDockerConfig().GetFlags(); c != "" {
		t.Errorf("service must be empty, got %q\n", c)
	}
	s.GetDockerConfig().Detach = false
	s.GetDockerConfig().Flags = ""
	if c := s.GetDockerConfig().GetFlags(); c != "-ti" {
		t.Errorf("service must have default flags %q,\n got %q\n", "-ti", c)
	}
}

func TestGetDetached(t *testing.T) {
	s := getService("service-with-pull-command", []string{"-v"})
	if c := s.GetDockerConfig().GetDetached(); c != "" {
		t.Error("when service has pull command, detach must be empty even if specified")
	}
	s.GetDockerConfig().Command = ""
	if c := s.GetDockerConfig().GetDetached(); c != "-d" {
		t.Errorf("service must have flag %q,\n got %q\n", "-d", c)
	}
	s.GetDockerConfig().Detach = false
	if c := s.GetDockerConfig().GetDetached(); c != "" {
		t.Errorf("service detach flag must be empty, got %q\n", c)
	}
}

func TestGetDockerCommand(t *testing.T) {
	s := getService("service-with-pull-command", []string{"-v"})
	if c := s.GetDockerConfig().GetDockerCommand(); c != "pull" {
		t.Errorf("wrong service docker command, expected %q\n got %q\n", "pull", c)
	}

	s.GetDockerConfig().Command = ""
	if c := s.GetDockerConfig().GetDockerCommand(); c != "run" {
		t.Errorf("wrong service docker command, must be empty got %q\n", "run")
	}
}

func TestGetContainerName(t *testing.T) {
	s := getService("service-with-conatiner-name", []string{"-v"})
	cn := s.GetDockerConfig().GetContainerName()
	e := "--name my_container"
	if e != cn {
		t.Errorf("name of cpntainer must be %q, got %q", e, cn)
	}

	s.GetDockerConfig().ContainerName = ""
	cn = s.GetDockerConfig().GetContainerName()
	e = ""
	if e != cn {
		t.Errorf("name of cpntainer must be %q, got %q", e, cn)
	}
}

func TestGetCleanUp(t *testing.T) {
	s := getService("service-with-clean-up", []string{"-v"})
	if c := s.GetDockerConfig().GetCleanUp(); c != "--rm" {
		t.Errorf("wrong service clean up param, expected %q\n got %q\n", "--rm", c)
	}

	s.GetDockerConfig().CleanUp = false
	if c := s.GetDockerConfig().GetCleanUp(); c != "" {
		t.Errorf("wrong service clean up param, must be empty got %q\n", c)
	}
}

func TestGetImageCommand(t *testing.T) {
	s := getService("service", []string{"-v"})
	if c := s.GetImageCommand(); c != "service" {
		t.Errorf("wrong service client command, expected %q got %q\n", "service", c)
	}

	s.GetDockerConfig().PreCommands = []string{"test"}
	if c := s.GetImageCommand(); c != "/bin/bash -c" {
		t.Errorf("wrong service client command, expected %q got %q\n", "/bin/bash -c", c)
	}

	s.GetDockerConfig().PreCommands = nil
	s.GetDockerConfig().PostCommands = []string{"test"}
	if c := s.GetImageCommand(); c != "/bin/bash -c" {
		t.Errorf("wrong service client command, expected %q got %q\n", "/bin/bash -c", c)
	}

	s.GetDockerConfig().BinBash = true
	s.GetDockerConfig().PostCommands = nil
	if c := s.GetImageCommand(); c != "/bin/bash -c service" {
		t.Errorf("wrong service client command, expected %q got %q\n", "/bin/bash -c service", c)
	}

	s.GetDockerConfig().ImageCommand = ""
	if c := s.GetImageCommand(); c != "" {
		t.Errorf("wrong service client command, expected %q got %q\n", "empty string", c)
	}
}

func TestGetDockerConfig(t *testing.T) {
	getServiceDockerConfig(t)
}

func getServiceDockerConfig(t *testing.T) {
	s := getService("service", []string{"-v"})

	if c := s.GetDockerConfig(); c == nil {
		t.Errorf("docker config for service incorrect")
	}

	s.Config = nil
	if c := s.GetDockerConfig(); c != nil {
		t.Errorf("docker config for service must be empty")
	}
}

func TestGetWorkDir(t *testing.T) {
	s := getService("service", []string{})
	if wd := s.GetWorkDir(); wd != "" {
		t.Errorf("service must have empty work-dir, got: %q", wd)
	}

	e := "--workdir=/some/work/dir"
	s.Config.WorkDir = "/some/work/dir"
	if wd := s.GetWorkDir(); wd != e {
		t.Errorf("service must have work-dir %q, got: %q", e, wd)
	}

	e = "--workdir=/some/another/dir"
	s.Config.WorkDir = ""
	s.Config.HomeDir = "/some/another/dir"
	if wd := s.GetWorkDir(); wd != e {
		t.Errorf("service must have work-dir %q, got: %q", e, wd)
	}
}

func getService(n string, args []string) *Service {
	c := setConfig()
	s := c.GetService(n)
	return NewService(s, args)
}

func setConfig() *config.Config {
	config.NewConfig(libs.GetPwd() + "/testdata/satellite")
	c := config.GetConfig()
	return c
}
