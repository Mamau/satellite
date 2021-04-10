package services

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mamau/starter/config/docker"

	"github.com/mamau/starter/config"
	"github.com/mamau/starter/entity"
	"github.com/mamau/starter/libs"
)

func TestDockerConfigCommand(t *testing.T) {
	getComposerDockerConfigCommand(t)
	getYarnDockerConfigCommand(t)
	getBowerDockerConfigCommand(t)
}

func TestClientCommand(t *testing.T) {
	getComposerClientCommand(t)
	getYarnClientCommand(t)
	getBowerClientCommand(t)
}

func TestCollectCommand(t *testing.T) {
	getComposerCollectCommand(t)
	getYarnCollectCommand(t)
	getBowerCollectCommand(t)
}

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

func getBowerCollectCommand(t *testing.T) {
	bower := getBower()
	c := NewCollector(bower)
	e := fmt.Sprintf("-u 501 --workdir=/home/node %s mamau/bower some bower command; bower install; some bower post cmd", bower.GetProjectVolume())
	if cc := c.CollectCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong bower collect command, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}

	bower.Config.Docker = docker.Docker{}
	c = NewCollector(bower)
	e = fmt.Sprintf("--workdir=/home/node %s mamau/bower bower install", bower.GetProjectVolume())
	if cc := c.CollectCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong bower collect command with docker empty config, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}
}

func getYarnCollectCommand(t *testing.T) {
	yarn := getYarn()
	c := NewCollector(yarn)
	e := fmt.Sprintf("-u 501 -e SOME_VAR=someVal --add-host=host.docker.internal:127.0.0.1 -p 127.0.0.1:443:443 -p 127.0.0.1:80:80 -p 8080:8080 --dns=8.8.8.8 --dns=8.8.4.4 --workdir=/home/node -v /Users/mamau/go/src/github.com/mamau/starter/cache:/tmp -v /Users/mamau/go/src/github.com/mamau/starter:/image/volume %s node:10 /bin/bash -c yarn config set strict-ssl false --global; yarn config set version-tag-prefix v --global; yarn config set version-git-tag true --global; yarn config set version-commit-hooks true --global; yarn config set version-git-sign false --global; yarn config set bin-links true --global; yarn config set ignore-scripts false --global; yarn config set ignore-optional false --global; yarn config set strict-ssl false; npm config set; yarn install; npm config set; npm config second post cmd", yarn.GetProjectVolume())
	if cc := c.CollectCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong yarn collect command, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}

	yarn.Config.Docker = docker.Docker{}
	c = NewCollector(yarn)
	e = fmt.Sprintf("--workdir=/home/node %s node:12 /bin/bash -c yarn config set strict-ssl false --global; yarn config set version-tag-prefix v --global; yarn config set version-git-tag true --global; yarn config set version-commit-hooks true --global; yarn config set version-git-sign false --global; yarn config set bin-links true --global; yarn config set ignore-scripts false --global; yarn config set ignore-optional false --global; yarn install", yarn.GetProjectVolume())
	if cc := c.CollectCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong yarn collect command with docker empty config, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}

	yarn.Config.Docker = docker.Docker{}
	yarn.Config.Config = nil
	c = NewCollector(yarn)
	e = fmt.Sprintf("--workdir=/home/node %s node:12 /bin/bash -c yarn install", yarn.GetProjectVolume())
	if cc := c.CollectCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong yarn collect command with docker empty and config empty, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}
}

func getComposerCollectCommand(t *testing.T) {
	composer := getComposer()
	c := NewCollector(composer)
	e := fmt.Sprintf("--workdir=/home/www-data -v /Users/mamau/go/src/github.com/mamau/starter/cache:/tmp -v /Users/mamau/go/src/github.com/mamau/starter:/image/volume -v /Users/mamau/go/src/github.com/mamau/starter2:/image/volume2 %s composer:2 /bin/bash -c composer config --global process-timeout 400; composer config --global http-basic.github.com mamau some-token; composer config --global http-basic.gitlab.com mamau some-token; composer config --global optimize-autoloader false; composer config set any; composer command; composer install --ignore-platform-reqs; composer post cmd; composer post cmd2", composer.GetProjectVolume())
	if cc := c.CollectCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong composer collect command, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}

	composer.Config.Docker = docker.Docker{}
	c = NewCollector(composer)
	e = fmt.Sprintf("--workdir=/home/www-data %s composer:1.9 /bin/bash -c composer config --global process-timeout 400; composer config --global http-basic.github.com mamau some-token; composer config --global http-basic.gitlab.com mamau some-token; composer config --global optimize-autoloader false; composer install --ignore-platform-reqs", composer.GetProjectVolume())
	if cc := c.CollectCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong composer collect command with docker empty config, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}

	composer.Config.Docker = docker.Docker{}
	composer.Config.Config = nil
	c = NewCollector(composer)
	e = fmt.Sprintf("--workdir=/home/www-data %s composer:1.9 /bin/bash -c composer install --ignore-platform-reqs", composer.GetProjectVolume())
	if cc := c.CollectCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong composer collect command with docker empty and config empty, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}
}

func getBowerClientCommand(t *testing.T) {
	bower := getBower()
	c := NewCollector(bower)
	e := "some bower command; bower install; some bower post cmd"
	if cc := c.ClientCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong bower client command, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}

	bower.Config.Docker = docker.Docker{}
	c = NewCollector(bower)
	e = "bower install"
	if cc := c.ClientCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong client cmd when docker empty \nexpected: %q \n got: %q", e, strings.Join(cc, " "))
	}
}

func getYarnClientCommand(t *testing.T) {
	yarn := getYarn()
	c := NewCollector(yarn)
	e := "yarn config set strict-ssl false; npm config set; yarn install; npm config set; npm config second post cmd"
	if cc := c.ClientCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong yarn client command, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}

	yarn.Config.Docker = docker.Docker{}
	c = NewCollector(yarn)
	e = "yarn install"
	if cc := c.ClientCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong client cmd when docker empty \nexpected: %q \n got: %q", e, strings.Join(cc, " "))
	}
}

func getComposerClientCommand(t *testing.T) {
	composer := getComposer()
	c := NewCollector(composer)
	e := "composer config set any; composer command; composer install --ignore-platform-reqs; composer post cmd; composer post cmd2"
	if cc := c.ClientCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong composer client command, expected: %q \n got: %q \n", e, strings.Join(cc, " "))
	}

	composer.Config.Docker = docker.Docker{}
	c = NewCollector(composer)
	e = "composer install --ignore-platform-reqs"
	if cc := c.ClientCommand(); strings.Join(cc, " ") != e {
		t.Errorf("wrong client cmd when docker empty \nexpected: %q \n got: %q", e, strings.Join(cc, " "))
	}
}

func getBowerDockerConfigCommand(t *testing.T) {
	bower := getBower()
	c := NewCollector(bower)
	e := fmt.Sprintf("-u 501 --workdir=/home/node %s mamau/bower", bower.GetProjectVolume())
	if dcc := c.DockerConfigCommand(); strings.Join(dcc, " ") != e {
		t.Errorf("wrong bower config command, expected: %q \n got: %q \n", e, strings.Join(dcc, " "))
	}

	bower.Config.Docker = docker.Docker{}
	c = NewCollector(bower)
	e = fmt.Sprintf("--workdir=/home/node %s mamau/bower", bower.GetProjectVolume())
	if dcc := c.DockerConfigCommand(); strings.Join(dcc, " ") != e {
		t.Errorf("wrong config for empty docker \nexpected: %q \n got: %q", e, strings.Join(dcc, " "))
	}
}

func getYarnDockerConfigCommand(t *testing.T) {
	yarn := getYarn()
	c := NewCollector(yarn)
	e := fmt.Sprintf("-u 501 -e SOME_VAR=someVal --add-host=host.docker.internal:127.0.0.1 -p 127.0.0.1:443:443 -p 127.0.0.1:80:80 -p 8080:8080 --dns=8.8.8.8 --dns=8.8.4.4 --workdir=/home/node -v /Users/mamau/go/src/github.com/mamau/starter/cache:/tmp -v /Users/mamau/go/src/github.com/mamau/starter:/image/volume %s node:10", yarn.GetProjectVolume())
	if dcc := c.DockerConfigCommand(); strings.Join(dcc, " ") != e {
		t.Errorf("wrong yarn config command, expected: %q \n got: %q \n", e, strings.Join(dcc, " "))
	}

	yarn.Config.Docker = docker.Docker{}
	c = NewCollector(yarn)
	e = fmt.Sprintf("--workdir=/home/node %s node:12", yarn.GetProjectVolume())
	if dcc := c.DockerConfigCommand(); strings.Join(dcc, " ") != e {
		t.Errorf("wrong config for empty docker \nexpected: %q \n got: %q", e, strings.Join(dcc, " "))
	}
}

func getComposerDockerConfigCommand(t *testing.T) {
	composer := getComposer()
	c := NewCollector(composer)
	e := fmt.Sprintf("--workdir=/home/www-data -v /Users/mamau/go/src/github.com/mamau/starter/cache:/tmp -v /Users/mamau/go/src/github.com/mamau/starter:/image/volume -v /Users/mamau/go/src/github.com/mamau/starter2:/image/volume2 %s composer:2", composer.GetProjectVolume())
	if dcc := c.DockerConfigCommand(); strings.Join(dcc, " ") != e {
		t.Errorf("wrong docker config command, expected: %q \n got: %q \n", e, strings.Join(dcc, " "))
	}

	composer.Config.Docker = docker.Docker{}
	c = NewCollector(composer)
	e = fmt.Sprintf("--workdir=/home/www-data %s composer:1.9", composer.GetProjectVolume())
	if dcc := c.DockerConfigCommand(); strings.Join(dcc, " ") != e {
		t.Errorf("wrong config for empty docker \nexpected: %q \n got: %q", e, strings.Join(dcc, " "))
	}
}

func getComposer() *entity.Composer {
	setConfig()
	return entity.NewComposer("1.9", []string{"install", "--ignore-platform-reqs"})
}

func getBower() *entity.Bower {
	setConfig()
	return entity.NewBower([]string{"install"})
}

func getYarn() *entity.Yarn {
	setConfig()
	return entity.NewYarn("12", []string{"install"})
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
