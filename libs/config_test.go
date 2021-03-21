package libs

import (
	"strings"
	"testing"
)

func TestGetConfig(t *testing.T) {
	nc := NewConfig()
	nc.Path = GetPwd() + "/starter_not_exists"
	c := GetConfig()
	if c != nc {
		t.Errorf("on not exists file config must be empty")
	}

	if c.GetComposer() != nil {
		t.Error("on empty file starter composer must be nil")
	}

	c.Path = GetPwd() + "/testdata/starter"
	c = GetConfig()
	if c.GetComposer() == nil {
		t.Error("composer is nil")
	}

	if c != nc {
		t.Errorf("on exists file config must be not empty")
	}

	cleanServices(nc, t)
}

func TestGetGulp(t *testing.T) {
	c := getConfig("/testdata/starter")
	gulp := c.GetGulp()
	if gulp == nil {
		t.Error("gulp is empty")
	}
	cleanServices(c, t)
}

func TestGetBower(t *testing.T) {
	c := getConfig("/testdata/starter")
	gulp := c.GetBower()
	if gulp == nil {
		t.Error("bower is empty")
	}
	cleanServices(c, t)
}

func TestGetComposer(t *testing.T) {
	c := getConfig("/testdata/starter")
	gulp := c.GetComposer()
	if gulp == nil {
		t.Error("composer is empty")
	}
	cleanServices(c, t)
}

func TestGetYarn(t *testing.T) {
	c := getConfig("/testdata/starter")
	gulp := c.GetYarn()
	if gulp == nil {
		t.Error("yarn is empty")
	}
	cleanServices(c, t)
}

func TestGetPreCommands(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	if c.GetComposer().GetPreCommands() != "composer some cmd; composer some cmd2" {
		t.Error("pre command is not match")
	}

	cleanServices(c, t)
}

func TestGetPostCommands(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	if c.GetComposer().GetPostCommands() != "composer some post cmd" {
		t.Error("post command is not match")
	}

	cleanServices(c, t)
}

func TestGetCacheDir(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	if c.GetComposer().GetCacheDir() != "/Users/mamau/go/src/github.com/mamau/starter/cache" {
		t.Error("cache dir is not match")
	}
	cleanServices(c, t)
}

func TestGetWorkDir(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	if c.GetComposer().GetWorkDir() != "/home/www-data" {
		t.Error("work dir is not match")
	}
	cleanServices(c, t)
}

func TestGetUserId(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	uid := strings.Join(c.GetComposer().GetUserId(), " ")
	if uid != "-u 501" {
		t.Error("user id is not match")
	}
	cleanServices(c, t)
}

func TestGetEnvironmentVariables(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	e := strings.Join(c.GetComposer().GetEnvironmentVariables(), "; ")
	if e != "-e SOME_VAR=someVal" {
		t.Error("env vars is not match")
	}
	cleanServices(c, t)
}

func TestGetVersion(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	v := c.GetComposer().GetVersion()
	if v != "2" {
		t.Error("version is not match")
	}
	cleanServices(c, t)
}

func TestGetHosts(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	h := strings.Join(c.GetComposer().GetHosts(), "; ")
	if h != "--add-host=host.docker.internal:127.0.0.1; --add-host=anotherHost" {
		t.Error("hosts is not match")
	}
	cleanServices(c, t)
}

func TestGetPorts(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	p := strings.Join(c.GetComposer().GetPorts(), "; ")
	if p != "-p 127.0.0.1:443:443; -p 127.0.0.1:80:80; -p 8080:8080" {
		t.Error("ports is not match")
	}
	cleanServices(c, t)
}

func TestGetVolumes(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	v := strings.Join(c.GetComposer().GetVolumes(), "; ")
	if v != "-v /Users/mamau/go/src/github.com/mamau/starter:/image/volume; -v /Users/mamau/go/src/github.com/mamau/starter2:/image/volume2" {
		t.Error("volumes is not match")
	}
	cleanServices(c, t)
}

func TestGetDns(t *testing.T) {
	c := getConfig("/testdata/starterComposer")
	d := strings.Join(c.GetComposer().GetDns(), "; ")
	if d != "--dns=8.8.8.8; --dns=8.8.4.4" {
		t.Error("dns is not match")
	}
	cleanServices(c, t)
}

func getConfig(cn string) *Config {
	c := NewConfig()
	c.Path = GetPwd() + cn
	return GetConfig()
}

func cleanServices(nc *Config, t *testing.T) {
	t.Cleanup(func() {
		nc.Services.Composer = nil
		nc.Services.Bower = nil
		nc.Services.Gulp = nil
		nc.Services.Yarn = nil
	})
}
