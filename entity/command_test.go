package entity

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mamau/starter/config"
	"github.com/mamau/starter/libs"
)

func TestGetCommand(t *testing.T) {
	getBowerCommand(t)
	getComposerCommand(t)
	getYarnCommand(t)
}

func TestGetImage(t *testing.T) {
	getYarnImage(t)
	getComposerImage(t)
	getBowerImage(t)
}

func TestGetPreCommands(t *testing.T) {
	getYarnPreCommands(t)
	getComposerPreCommands(t)
	getBowerPreCommands(t)
}

func TestGetPostCommands(t *testing.T) {
	getYarnPostCommands(t)
	getComposerPostCommands(t)
	getBowerPostCommands(t)
}

func TestProjectVolume(t *testing.T) {
	getComposerProjectVolume(t)
	getYarnProjectVolume(t)
	getBowerProjectVolume(t)
}

func TestDockerCommandData(t *testing.T) {
	getComposerDockerCommandData(t)
	getYarnDockerCommandData(t)
	getBowerDockerCommandData(t)
}

func getBowerDockerCommandData(t *testing.T) {
	b := getBower([]string{})
	b.Args = []string{"install"}
	b.HomeDir = "/home/node"
	needle := fmt.Sprintf("-u 501 --workdir=/home/node %s mamau/bower some bower command; bower install; some bower post cmd", b.GetProjectVolume())
	if dbd := strings.Join(b.CollectCommand(), " "); dbd != needle {
		t.Errorf("wrong full command for bower expect \n %q got \n %q \n", needle, dbd)
	}
}

func getYarnDockerCommandData(t *testing.T) {
	y := getYarn("", []string{})
	y.Args = []string{"install"}
	y.HomeDir = "/home/node"
	needle := fmt.Sprintf("-u 501 -e SOME_VAR=someVal --add-host=host.docker.internal:127.0.0.1 -p 127.0.0.1:443:443 -p 127.0.0.1:80:80 -p 8080:8080 --dns=8.8.8.8 --dns=8.8.4.4 --workdir=/home/node -v /Users/mamau/go/src/github.com/mamau/starter/cache:/tmp -v /Users/mamau/go/src/github.com/mamau/starter:/image/volume %s node:10 /bin/bash -c yarn config set strict-ssl false; npm config set; yarn install; npm config set; npm config second post cmd", y.GetProjectVolume())
	if dyd := strings.Join(y.CollectCommand(), " "); dyd != needle {
		t.Errorf("wrong full command for yarn expect \n %q got \n %q \n", needle, dyd)
	}
}

func getComposerDockerCommandData(t *testing.T) {
	c := getComposer("", []string{})
	c.Args = []string{"install --ignore-platform-reqs"}
	c.HomeDir = "/home/www-data"
	needle := fmt.Sprintf("--workdir=/home/www-data -v /Users/mamau/go/src/github.com/mamau/starter/cache:/tmp -v /Users/mamau/go/src/github.com/mamau/starter:/image/volume -v /Users/mamau/go/src/github.com/mamau/starter2:/image/volume2 %s composer:2 /bin/bash -c composer config --global process-timeout 400; composer config --global http-basic.github.com mamau some-token; composer config --global http-basic.gitlab.com mamau some-token; composer config --global optimize-autoloader false; composer config set any; composer command; composer install --ignore-platform-reqs; composer post cmd; composer post cmd2", c.GetProjectVolume())
	if dcd := strings.Join(c.CollectCommand(), " "); dcd != needle {
		t.Errorf("wrong full command for composer \n got %q \n need %q", dcd, needle)
	}
}

func getBowerProjectVolume(t *testing.T) {
	b := getBower([]string{})
	mv := fmt.Sprintf("-v %s:%s", libs.GetPwd(), b.HomeDir)
	if pv := b.GetProjectVolume(); pv != mv {
		t.Errorf("something wrong with yarn bower, got: %q", pv)
	}
}

func getYarnProjectVolume(t *testing.T) {
	y := getYarn("", []string{})
	mv := fmt.Sprintf("-v %s:%s", libs.GetPwd(), y.HomeDir)
	if pv := y.GetProjectVolume(); pv != mv {
		t.Errorf("something wrong with yarn volumes, got: %q", pv)
	}
}

func getComposerProjectVolume(t *testing.T) {
	c := getComposer("", []string{})
	mv := fmt.Sprintf("-v %s:%s", libs.GetPwd(), c.HomeDir)
	if pv := c.GetProjectVolume(); pv != mv {
		t.Errorf("something wrong with composer volumes, got: %q", pv)
	}
}

func getBowerPostCommands(t *testing.T) {
	b := getBower([]string{})
	expected := "some bower post cmd"
	if pc := b.Config.GetPostCommands(); pc != expected {
		t.Errorf("wrong post-command format for bower, got %q", pc)
	}

	b.Config.SetPostCommands([]string{})
	if pc := b.Config.GetPostCommands(); pc != "" {
		t.Errorf("bower must be empty post-command if config settings empty, got %q", pc)
	}
}

func getComposerPostCommands(t *testing.T) {
	c := getComposer("", []string{})
	expected := "composer post cmd; composer post cmd2"
	if pc := c.Config.GetPostCommands(); pc != expected {
		t.Errorf("wrong post-command format for composer, got %q", pc)
	}

	c.Config.SetPostCommands([]string{})
	if pc := c.Config.GetPostCommands(); pc != "" {
		t.Errorf("yarn must be empty post-command if config settings empty, got %q", pc)
	}
}

func getYarnPostCommands(t *testing.T) {
	y := getYarn("", []string{})
	expected := "npm config set; npm config second post cmd"
	if pc := y.Config.GetPostCommands(); pc != expected {
		t.Errorf("wrong post-command format, got %q", pc)
	}

	y.Config.SetPostCommands([]string{})
	if pc := y.Config.GetPostCommands(); pc != "" {
		t.Errorf("yarn must be empty post-command if config settings empty, got %q", pc)
	}
}

func getBowerPreCommands(t *testing.T) {
	b := getBower([]string{})
	expected := "some bower command"
	if pc := b.Config.GetPreCommands(); pc != expected {
		t.Errorf("pre-command for bower must be %q got %q", expected, pc)
	}

	b.Config.SetPreCommands([]string{})
	if pc := b.Config.GetPreCommands(); pc != "" {
		t.Errorf("bower must be empty pre-command if config settings empty")
	}
}

func getComposerPreCommands(t *testing.T) {
	c := getComposer("", []string{})
	expected := "composer config set any; composer command"
	if pc := c.Config.GetPreCommands(); pc != expected {
		t.Errorf("pre-command for composer must be %q got %q", expected, pc)
	}

	c.Config.SetPreCommands([]string{})
	if pc := c.Config.GetPreCommands(); pc != "" {
		t.Errorf("composer must be empty pre-command if config settings empty")
	}
}

func getYarnPreCommands(t *testing.T) {
	y := getYarn("", []string{})
	expected := "yarn config set strict-ssl false; npm config set"
	if pc := y.Config.GetPreCommands(); pc != expected {
		t.Errorf("pre-command must be %q got %q", expected, pc)
	}

	y.Config.SetPreCommands([]string{})
	if pc := y.Config.GetPreCommands(); pc != "" {
		t.Errorf("yarn must be empty pre-command if config settings empty")
	}
}

func getBowerImage(t *testing.T) {
	b := getBower([]string{})
	if i := b.GetImage(); i != "mamau/bower" {
		t.Errorf("yarn image name must be %q got: %s", "mamau/bower", i)
	}
}

func getYarnImage(t *testing.T) {
	y := getYarn("12", []string{})
	if i := y.GetImage(); i != "node:10" {
		t.Errorf("yarn image name must be %q, (priory config) got: %s", "node:10", i)
	}

	y.Config.SetVersion("")
	y.Version = ""
	if i := y.GetImage(); i != "node" {
		t.Errorf("yarn image name without version must be %q, got: %s", "node", i)
	}

	y.Version = "10"
	if i := y.GetImage(); i != "node:10" {
		t.Errorf("yarn image name must be %q, got: %s", "node:10", i)
	}
}

func getComposerImage(t *testing.T) {
	c := getComposer("1.9", []string{})
	if i := c.GetImage(); i != "composer:2" {
		t.Errorf("composer image name must be %q, (priory config) got: %s", "composer:2", i)
	}
	c.Config.SetVersion("")
	c.Version = ""
	if i := c.GetImage(); i != "composer" {
		t.Errorf("composer image name without version must be %q, got: %s", "composer", i)
	}
	c.Version = "1.9"
	if i := c.GetImage(); i != "composer:1.9" {
		t.Errorf("composer image name must be %q, got: %s", "composer:1.9", i)
	}
}

func getBowerCommand(t *testing.T) {
	b := getBower([]string{"--help", "--version"})
	cmd := b.GetClientCommand()
	if cmd != "bower --help --version" {
		t.Error("wrong command")
	}

	b.Args = []string{}
	if b := b.GetClientCommand(); b != "bower" {
		t.Error("empty bower command must have default command: --version")
	}
}

func getComposerCommand(t *testing.T) {
	c := getComposer("1.9", []string{"install", "--ignore-platform-reqs"})
	if cmd := c.GetClientCommand(); cmd != "composer install --ignore-platform-reqs" {
		t.Errorf("wrong composer command, got: %s", cmd)
	}

	c.Args = []string{}
	if cmd := c.GetClientCommand(); cmd != "composer" {
		t.Errorf("composer with empty args must have command name %q", "composer")
	}
}

func getYarnCommand(t *testing.T) {
	y := getYarn("10", []string{"install"})
	if cmd := y.GetClientCommand(); cmd != "yarn install" {
		t.Errorf("yarn must be %q, got %s", "yarn install", cmd)
	}

	y.Args = []string{}
	if cmd := y.GetClientCommand(); cmd != "yarn" {
		t.Errorf("yarn with empty args must have command name %q", "yarn")
	}
}

func getComposer(v string, args []string) *Composer {
	setConfig()
	return NewComposer(v, args)
}

func getBower(args []string) *Bower {
	setConfig()
	return NewBower(args)
}

func getYarn(v string, args []string) *Yarn {
	setConfig()
	return NewYarn(v, args)
}

func setConfig() {
	c := config.GetConfig()
	c.Path = libs.GetPwd() + "/testdata/starter"
}
