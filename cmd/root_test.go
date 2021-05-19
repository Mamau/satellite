package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/mamau/satellite/strategy"

	"github.com/mamau/satellite/config"
	"github.com/mamau/satellite/libs"
)

func TestGetRunnableCommand(t *testing.T) {
	os.Args = []string{"sat", "php", "-v"}
	if cmd := getRunnableCommand(); cmd != "php" {
		t.Errorf("runnable command must be %q", "php")
	}

	os.Args = []string{"sat"}
	if cmd := getRunnableCommand(); cmd != "" {
		t.Error("runnable command must be empty")
	}
}

func TestGetAvailableCommands(t *testing.T) {
	ac := getAvailableCommands()
	var availableCommands []string
	for _, v := range rootCmd.Commands() {
		availableCommands = append(availableCommands, v.Name())
	}

	if strings.Join(ac, " ") != strings.Join(availableCommands, " ") {
		t.Errorf("available command expected %q\n got %q\n", strings.Join(ac, " "), strings.Join(availableCommands, " "))
	}
}

func TestDetermineStrategy(t *testing.T) {
	c := setConfig().GetService("fresh-img")
	strat := determineStrategy(c, []string{})
	st := strat.GetContext().Value("type")
	if st != strategy.PullType {
		t.Errorf("worng value context, expected %q, got %q", strategy.PullType, st)
	}

	c = setConfig().GetService("my-image")
	strat = determineStrategy(c, []string{})
	st = strat.GetContext().Value("type")
	if st != strategy.DaemonType {
		t.Errorf("worng value context, expected %q, got %q", strategy.DaemonType, st)
	}

	c = setConfig().GetService("composer")
	strat = determineStrategy(c, []string{})
	st = strat.GetContext().Value("type")
	if st != strategy.RunType {
		t.Errorf("worng value context, expected %q, got %q", strategy.RunType, st)
	}
}

func setConfig() *config.Config {
	config.NewConfig(libs.GetPwd() + "/testdata/satellite")
	c := config.GetConfig()
	return c
}
