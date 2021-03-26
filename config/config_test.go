package config

import (
	"testing"

	"github.com/mamau/starter/config/composer"

	"github.com/mamau/starter/libs"
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
	cr := c.GetComposer()
	if cr == nil {
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

func TestGetProcessTimeoutAsCommand(t *testing.T) {
	c := getComposerConfig()
	pt := c.GetProcessTimeoutAsCommand()
	if pt != "composer config --global process-timeout 400" {
		t.Errorf("wrong proccess-timeout value, got: %s", pt)
	}

	c.ProcessTimeout = ""
	if pt := c.GetProcessTimeoutAsCommand(); pt != "" {
		t.Errorf("when ProcessTimeout is empty, return value must be is empty string")
	}
}

func TestGetRepoAsCommand(t *testing.T) {
	c := getComposerConfig()
	r := c.GetRepoAsCommand()
	if r != "composer config --global http-basic.github.com mamau some-token; composer config --global http-basic.gitlab.com mamau some-token" {
		t.Error("wrong repositories format")
	}
	c.Repositories = nil
	if r := c.GetRepoAsCommand(); r != "" {
		t.Error("when repositories is empty, value must be empty string")
	}
}

func TestGetOptimizeAutoloaderAsCommand(t *testing.T) {
	c := getComposerConfig()
	oa := c.GetOptimizeAutoloaderAsCommand()
	if oa != "composer config --global optimize-autoloader false" {
		t.Error("wrong optimize-autoloader format")
	}
	c.OptimizeAutoloader = ""
	if oa := c.GetOptimizeAutoloaderAsCommand(); oa != "" {
		t.Error("when optimize-autoloader is empty, value must be empty string")
	}
}

func getComposerConfig() *composer.Config {
	c := getConfig("/testdata/composer_config")
	return c.GetComposer().Config
}

func TestGetPreCommands(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	if c.GetComposer().GetPreCommands() != "composer some cmd; composer some cmd2" {
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
	expected := "-v /Users/mamau/go/src/github.com/mamau/starter/cache:/tmp"
	if v := c.GetComposer().GetCacheVolume(); v != expected {
		t.Errorf("cache dir must bet %q, got %q", expected, v)
	}
	cleanServices(c, t)
}

func TestGetWorkDir(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	if c.GetComposer().GetWorkDir() != "--workdir=/home/www-data" {
		t.Errorf("work dir must be %q, got %q", "--workdir=/home/www-data", c.GetComposer().GetWorkDir())
	}
	cleanServices(c, t)
}

func TestGetUserId(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	expected := "-u 501"
	if uid := c.GetComposer().GetUserId(); uid != expected {
		t.Errorf("user id must be %q, got %q", expected, uid)
	}
	cleanServices(c, t)
}

func TestGetEnvironmentVariables(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	expected := "-e SOME_VAR=someVal"
	if e := c.GetComposer().GetEnvironmentVariables(); e != expected {
		t.Errorf("env vars must be %q, got %q", expected, e)
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
	expected := "--add-host=host.docker.internal:127.0.0.1 --add-host=anotherHost"
	if h := c.GetComposer().GetHosts(); h != expected {
		t.Errorf("hosts must be %q, got %q", expected, h)
	}
	cleanServices(c, t)
}

func TestGetPorts(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	expected := "-p 127.0.0.1:443:443 -p 127.0.0.1:80:80 -p 8080:8080"
	if p := c.GetComposer().GetPorts(); p != expected {
		t.Errorf("ports must be %q, got %q", expected, p)
	}
	cleanServices(c, t)
}

func TestGetVolumes(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	expected := "-v /Users/mamau/go/src/github.com/mamau/starter:/image/volume -v /Users/mamau/go/src/github.com/mamau/starter2:/image/volume2"
	if v := c.GetComposer().GetVolumes(); v != expected {
		t.Errorf("volumes must be %q, got %q", expected, v)
	}
	cleanServices(c, t)
}

func TestGetDns(t *testing.T) {
	c := getConfig("/testdata/starter_composer")
	expected := "--dns=8.8.8.8 --dns=8.8.4.4"
	if d := c.GetComposer().GetDns(); d != expected {
		t.Errorf("dns must be %q, got %q", expected, d)
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
