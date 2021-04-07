package entity

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mamau/starter/config/yarn"

	"github.com/mamau/starter/config/composer"

	"github.com/mamau/starter/config/docker"

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

func TestGetClientSignature(t *testing.T) {
	getBowerClientSignature(t)
	getComposerClientSignature(t)
	getYarnClientSignature(t)
	getServiceClientSignature(t)
}

func TestGetClientCommand(t *testing.T) {
	s := getService("php", []string{"-v"})
	if c := s.GetDockerConfig().GetClientCommand(); c != "php" {
		t.Errorf("wrong service client command, expected %q\n got %q\n", "php", c)
	}
}

func getServiceClientSignature(t *testing.T) {
	data := []string{"some", "data"}
	s := getService("php", []string{"-v"})
	r := s.GetClientSignature(data)
	e := "some data"
	if e != strings.Join(r, " ") {
		t.Errorf("service client signature must be %q\n got %q", e, strings.Join(r, " "))
	}
}

func TestConfigToCommand(t *testing.T) {
	getComposerConfigToCommand(t)
	getYarnConfigToCommand(t)
}

func TestNewService(t *testing.T) {
	service := NewService(nil, []string{})
	if service != nil {
		t.Error("service must be nil")
	}
}

func getYarnConfigToCommand(t *testing.T) {
	y := getYarn("", []string{})
	data := y.configToCommand()
	e := "yarn config set strict-ssl false --global; yarn config set version-tag-prefix v --global; yarn config set version-git-tag true --global; yarn config set version-commit-hooks true --global; yarn config set version-git-sign false --global; yarn config set bin-links true --global; yarn config set ignore-scripts false --global; yarn config set ignore-optional false --global;"
	if e != strings.Join(data, " ") {
		t.Errorf("error yarn config to config, expect %q\n got %q\n", e, strings.Join(data, " "))
	}
	y.Config.Config = nil
	if d := y.configToCommand(); len(d) != 0 {
		t.Errorf("config must be empty got %q", d)
	}

	y.Config.Config = &yarn.Config{}
	if d := y.configToCommand(); len(d) != 0 {
		t.Errorf("config must be empty got %q", d)
	}

	y.Config = nil
	if d := y.configToCommand(); len(d) != 0 {
		t.Errorf("config must be empty got %q", d)
	}
}

func getComposerConfigToCommand(t *testing.T) {
	c := getComposer("", []string{})
	data := c.configToCommand()
	e := "composer config --global process-timeout 400; composer config --global http-basic.github.com mamau some-token; composer config --global http-basic.gitlab.com mamau some-token; composer config --global optimize-autoloader false;"
	if e != strings.Join(data, " ") {
		t.Errorf("error composer config to config, expect %q\n got %q\n", e, strings.Join(data, " "))
	}
	c.Config.Config = nil
	if d := c.configToCommand(); len(d) != 0 {
		t.Errorf("config must be empty got %q", d)
	}

	c.Config.Config = &composer.Config{}
	if d := c.configToCommand(); len(d) != 0 {
		t.Errorf("config must be empty got %q", d)
	}

	c.Config = nil
	if d := c.configToCommand(); len(d) != 0 {
		t.Errorf("config must be empty got %q", d)
	}
}

func getYarnClientSignature(t *testing.T) {
	data := []string{"some", "data"}
	y := getYarn("", []string{})
	r := y.GetClientSignature(data)
	e := "/bin/bash -c yarn config set strict-ssl false --global; yarn config set version-tag-prefix v --global; yarn config set version-git-tag true --global; yarn config set version-commit-hooks true --global; yarn config set version-git-sign false --global; yarn config set bin-links true --global; yarn config set ignore-scripts false --global; yarn config set ignore-optional false --global; some data"
	if e != strings.Join(r, " ") {
		t.Errorf("yarn client signature must be %q\n got %q", e, strings.Join(r, " "))
	}
}

func getComposerClientSignature(t *testing.T) {
	data := []string{"some", "data"}
	c := getComposer("", []string{})
	r := c.GetClientSignature(data)
	e := "/bin/bash -c composer config --global process-timeout 400; composer config --global http-basic.github.com mamau some-token; composer config --global http-basic.gitlab.com mamau some-token; composer config --global optimize-autoloader false; some data"
	if e != strings.Join(r, " ") {
		t.Errorf("composer client signature must be %q\n got %q", e, strings.Join(r, " "))
	}
}

func getBowerClientSignature(t *testing.T) {
	data := []string{"some", "data"}
	b := getBower([]string{})
	c := b.GetClientSignature(data)
	e := "some data"
	if e != strings.Join(c, " ") {
		t.Errorf("bower client signature must be %q\n got %q", e, strings.Join(c, " "))
	}
}

func TestGetCommandConfig(t *testing.T) {
	getBowerCommandConfig(t)
	getComposerCommandConfig(t)
	getYarnCommandConfig(t)
	getServiceCommandConfig(t)
}

func getServiceCommandConfig(t *testing.T) {
	s := getService("php", []string{"-v"})
	c := s.GetCommandConfig()
	if c == nil {
		t.Error("command service cannot be empty")
	}
}

func getYarnCommandConfig(t *testing.T) {
	y := getYarn("", []string{})
	c := y.GetCommandConfig()
	if c == nil {
		t.Error("command yarn cannot be empty")
	}
}

func getComposerCommandConfig(t *testing.T) {
	c := getComposer("", []string{})
	cc := c.GetCommandConfig()
	if cc == nil {
		t.Error("command composer cannot be empty")
	}
}

func getBowerCommandConfig(t *testing.T) {
	b := getBower([]string{})
	c := b.GetCommandConfig()
	if c == nil {
		t.Error("command bower cannot be empty")
	}
}

func TestGetDockerConfig(t *testing.T) {
	getBowerDockerConfig(t)
	getComposerDockerConfig(t)
	getYarnDockerConfig(t)
	getServiceDockerConfig(t)
}

func getServiceDockerConfig(t *testing.T) {
	s := getService("php", []string{"-v"})

	if c := s.GetDockerConfig(); c == nil {
		t.Errorf("docker config for service incorrect")
	}

	s.Config = nil
	if c := s.GetDockerConfig(); c != nil {
		t.Errorf("docker config for service must be empty")
	}
}

func getYarnDockerConfig(t *testing.T) {
	y := getYarn("", []string{})
	c := y.GetDockerConfig()
	if c == nil {
		t.Errorf("docker config for yarn incorrect")
	}
	y.Config = nil
	c = y.GetDockerConfig()
	if c != nil {
		t.Errorf("docker config for yarn incorrect when config is empty")
	}
}

func getComposerDockerConfig(t *testing.T) {
	c := getComposer("", []string{})
	cc := c.GetDockerConfig()
	if cc == nil {
		t.Errorf("docker config for composer incorrect")
	}

	c.Config.Docker = docker.Docker{}
	cc = c.GetDockerConfig()
	if cc != nil {
		t.Errorf("docker config for composer incorrect when config docker is empty")
	}

	c.Config = nil
	cc = c.GetDockerConfig()
	if cc != nil {
		t.Errorf("docker config for composer incorrect when config is empty")
	}
}

func getBowerDockerConfig(t *testing.T) {
	b := getBower([]string{})
	c := b.GetDockerConfig()
	if c == nil {
		t.Errorf("docker config for bower incorrect")
	}
	b.Config = nil
	c = b.GetDockerConfig()
	if c != nil {
		t.Errorf("docker config for bower incorrect when config is empty")
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
	y.Config = nil
	if i := y.GetImage(); i != "node:10" {
		t.Errorf("yarn image name must be %q, got: %s", "composer:1.9", i)
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

	c.Config = nil
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

	b.CmdName = ""
	if b := b.GetClientCommand(); b != "--help --version" {
		t.Error("error client command when empty CmdName")
	}

	b.CmdName = "bower"
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

	c.CmdName = ""
	if b := c.GetClientCommand(); b != "install --ignore-platform-reqs" {
		t.Error("error client command when empty CmdName")
	}

	c.CmdName = "composer"
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

	y.CmdName = ""
	if b := y.GetClientCommand(); b != "install" {
		t.Error("error client command when empty CmdName")
	}

	y.CmdName = "yarn"
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

func getService(n string, args []string) *Service {
	c := setConfig()
	fmt.Println(c)
	s := c.GetService(n)
	return NewService(s, args)
}

func setConfig() *config.Config {
	config.NewConfig(libs.GetPwd() + "/testdata/starter")
	c := config.GetConfig()
	return c
}
