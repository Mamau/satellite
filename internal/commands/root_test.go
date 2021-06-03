package commands

import (
	"fmt"
	"testing"

	"github.com/mamau/satellite/internal/strategy"

	"github.com/mamau/satellite/pkg"

	"github.com/mamau/satellite/internal/config"
)

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
	config.NewConfig(fmt.Sprintf("%s/testdata/satellite", pkg.GetPwd()))
	c := config.GetConfig()
	return c
}
