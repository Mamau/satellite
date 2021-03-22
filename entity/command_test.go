package entity

import (
	"testing"

	"github.com/mamau/starter/config"
	"github.com/mamau/starter/libs"
)

func TestGetCommand(t *testing.T) {
	args := []string{"--help", "--version"}
	b := getBower(args)
	cmd := b.getCommand()
	if cmd != "--help --version" {
		t.Error("wrong command")
	}

	b.Args = []string{}
	if b := b.getCommand(); b != "" {
		t.Error("empty bower command must have default command: --version")
	}

	c := getComposer("1.9", []string{"install", "--ignore-platform-reqs"})
	if cmd = c.getCommand(); cmd != "composer install --ignore-platform-reqs" {
		t.Errorf("wrong composer command, got: %s", cmd)
	}

	c.Args = []string{}
	if cmd := c.getCommand(); cmd != "composer" {
		t.Error("composer with empty args must have command name `composer`")
	}

}

func TestGetImage(t *testing.T) {

}

func getComposer(v string, args []string) *Composer {
	setConfig()
	return NewComposer(v, args)
}

func getBower(args []string) *Bower {
	setConfig()
	return NewBower(args)
}

func setConfig() {
	c := config.GetConfig()
	c.Path = libs.GetPwd() + "/testdata/starter"
}
