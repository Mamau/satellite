package config

import (
	"github.com/mamau/starter/libs"
	"strings"
	"testing"
)

func TestGetConfig(t *testing.T) {
	nc := NewConfig()
	nc.Path = libs.GetPwd() + "/starter_not_exists"
	c := GetConfig()
	if c != nc {
		t.Errorf("on not exists file config must be empty")
	}

	if c.GetComposer() != nil {
		t.Error("on empty file starter composer must be nil")
	}

	c.Path = libs.GetPwd() + "/testdata/starter"
	c = GetConfig()
	if c.GetComposer() == nil {
		t.Error("composer is nil")
	}

	if c != nc {
		t.Errorf("on exists file config must be not empty")
	}

	cleanServices(nc, t)
}

func TestGetBower(t *testing.T) {
	c := getConfig("/testdata/starter")
	bower := c.GetBower()
	if bower == nil {
		t.Error("bower is empty")
	}
	cleanServices(c, t)
}

func TestGetComposer(t *testing.T) {
	c := getConfig("/testdata/starter")
	composer := c.GetComposer()
	if composer == nil {
		t.Error("composer is empty")
	}
	cleanServices(c, t)
}

func TestGetYarn(t *testing.T) {
	c := getConfig("/testdata/starter")
	yarn := c.GetYarn()
	if yarn == nil {
		t.Error("yarn is empty")
	}
	cleanServices(c, t)
}

func TestGetPreCommands(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	if c.GetComposer().GetPreCommands() != "composer some cmd; composer some cmd2; " {
		t.Error("pre command is not match")
	}

	cleanServices(c, t)
}

func TestGetPostCommands(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	if c.GetComposer().GetPostCommands() != "composer some post cmd" {
		t.Error("post command is not match")
	}

	cleanServices(c, t)
}

func TestGetCacheDir(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	if c.GetComposer().GetCacheDir() != "/Users/mamau/go/src/github.com/mamau/starter/cache" {
		t.Error("cache dir is not match")
	}
	cleanServices(c, t)
}

func TestGetWorkDir(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	if c.GetComposer().GetWorkDir() != "/home/www-data" {
		t.Error("work dir is not match")
	}
	cleanServices(c, t)
}

func TestGetUserId(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	uid := strings.Join(c.GetComposer().GetUserId(), " ")
	if uid != "-u 501" {
		t.Error("user id is not match")
	}
	cleanServices(c, t)
}

func TestGetEnvironmentVariables(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	e := strings.Join(c.GetComposer().GetEnvironmentVariables(), "; ")
	if e != "-e SOME_VAR=someVal" {
		t.Error("env vars is not match")
	}
	cleanServices(c, t)
}

func TestGetVersion(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	v := c.GetComposer().GetVersion()
	if v != "2" {
		t.Error("version is not match")
	}
	cleanServices(c, t)
}

func TestGetHosts(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	h := strings.Join(c.GetComposer().GetHosts(), "; ")
	if h != "--add-host=host.docker.internal:127.0.0.1; --add-host=anotherHost" {
		t.Error("hosts is not match")
	}
	cleanServices(c, t)
}

func TestGetPorts(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	p := strings.Join(c.GetComposer().GetPorts(), "; ")
	if p != "-p 127.0.0.1:443:443; -p 127.0.0.1:80:80; -p 8080:8080" {
		t.Error("ports is not match")
	}
	cleanServices(c, t)
}

func TestGetVolumes(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	v := strings.Join(c.GetComposer().GetVolumes(), "; ")
	if v != "-v /Users/mamau/go/src/github.com/mamau/starter:/image/volume; -v /Users/mamau/go/src/github.com/mamau/starter2:/image/volume2" {
		t.Error("volumes is not match")
	}
	cleanServices(c, t)
}

func TestGetDns(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	d := strings.Join(c.GetComposer().GetDns(), "; ")
	if d != "--dns=8.8.8.8; --dns=8.8.4.4" {
		t.Error("dns is not match")
	}
	cleanServices(c, t)
}

func TestGetClientConfig(t *testing.T) {
	fp := libs.GetPwd() + "/testdata/starter"
	result := GetClientConfig(fp)

	if result != fp+".yaml" {
		t.Errorf("file %s is not exist", fp)
	}

	fp = libs.GetPwd() + "/testdata/starter_not_exists"
	result = GetClientConfig(fp)
	if result != "" {
		t.Errorf("file %s not exists and return non empty string", fp)
	}
}

func getConfig(cn string) *Config {
	c := NewConfig()
	c.Path = libs.GetPwd() + cn
	return GetConfig()
}

func cleanServices(nc *Config, t *testing.T) {
	t.Cleanup(func() {
		nc.Commands.Composer = nil
		nc.Commands.Bower = nil
		nc.Commands.Yarn = nil
	})
}
