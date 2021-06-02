package commands

import (
	"os"
	"strings"
	"testing"

	"github.com/mamau/satellite/pkg"

	config2 "github.com/mamau/satellite/internal/config"

	strategy2 "github.com/mamau/satellite/internal/strategy"
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
	if st != strategy2.PullType {
		t.Errorf("worng value context, expected %q, got %q", strategy2.PullType, st)
	}

	c = setConfig().GetService("my-image")
	strat = determineStrategy(c, []string{})
	st = strat.GetContext().Value("type")
	if st != strategy2.DaemonType {
		t.Errorf("worng value context, expected %q, got %q", strategy2.DaemonType, st)
	}

	c = setConfig().GetService("composer")
	strat = determineStrategy(c, []string{})
	st = strat.GetContext().Value("type")
	if st != strategy2.RunType {
		t.Errorf("worng value context, expected %q, got %q", strategy2.RunType, st)
	}
}

func setConfig() *config2.Config {
	config2.NewConfig(pkg.GetPwd() + "/testdata/satellite")
	c := config2.GetConfig()
	return c
}
