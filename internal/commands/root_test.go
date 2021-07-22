package commands

import (
	"fmt"
	"testing"

	"github.com/mamau/satellite/internal/config/docker"

	"github.com/mamau/satellite/pkg"

	"github.com/mamau/satellite/internal/config"
)

func TestValidation(t *testing.T) {
	c := setConfig().GetService("fresh-img")
	strat := determineStrategy(c, []string{})
	err := validation(strat)
	if err != nil {
		t.Errorf("error must be nil on %q service\n", strat.GetContext().GetConfig().GetType())
	}

	c = setConfig().GetService("my-image")
	strat = determineStrategy(c, []string{})
	err = validation(strat)
	if err != nil {
		t.Errorf("error must be nil on %q service\n", strat.GetContext().GetConfig().GetType())
	}

	c = setConfig().GetService("composer")
	strat = determineStrategy(c, []string{})
	err = validation(strat)
	if err == nil {
		t.Errorf("error cannot be nil when empty arguments for service %q\n", strat.GetContext().GetConfig().GetType())
	}
}

func TestDetermineStrategy(t *testing.T) {
	c := setConfig().GetService("run-docker-compose")
	strat := determineStrategy(c, []string{})
	st := strat.GetContext()
	r := st.GetConfig().GetType()

	if r != docker.DOCKER_COMPOSE {
		t.Errorf("worng value context, expected %q, got %q", docker.DOCKER_COMPOSE, r)
	}

	c = setConfig().GetService("fresh-img")
	strat = determineStrategy(c, []string{})
	st = strat.GetContext()
	r = st.GetConfig().GetType()

	if r != docker.PULL {
		t.Errorf("worng value context, expected %q, got %q", docker.PULL, r)
	}

	c = setConfig().GetService("my-image")
	strat = determineStrategy(c, []string{})
	st = strat.GetContext()
	r = st.GetConfig().GetType()
	if r != docker.DAEMON {
		t.Errorf("worng value context, expected %q, got %q", docker.DAEMON, r)
	}

	c = setConfig().GetService("composer")
	strat = determineStrategy(c, []string{})
	st = strat.GetContext()
	r = st.GetConfig().GetType()
	if r != docker.RUN {
		t.Errorf("worng value context, expected %q, got %q", docker.RUN, r)
	}
}

func setConfig() *config.Config {
	config.NewConfig(fmt.Sprintf("%s/testdata/satellite", pkg.GetPwd()))
	c := config.GetConfig()
	return c
}
