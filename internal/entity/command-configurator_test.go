package entity

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientCommand(t *testing.T) {
	run := Run{}
	cc := newConfigConfigurator([]string{}, []string{}, &run)
	result := strings.Join(cc.getClientCommand(), " ")
	assert.Empty(t, result)

	cc = newConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result = strings.Join(cc.getClientCommand(), " ")
	assert.Equal(t, result, "ls -la")

	run.PreCommands = []string{"some command"}
	cc = newConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result = strings.Join(cc.getClientCommand(), " ")
	assert.Equal(t, result, "/bin/bash -c some command; ls -la")
}

func TestPrepareCommand(t *testing.T) {
	run := Run{}
	cc := newConfigConfigurator([]string{}, []string{}, &run)
	assert.Empty(t, cc.prepareCommand())

	run.PreCommands = []string{"some command"}
	cc = newConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result := strings.Join(cc.prepareCommand(), " ")
	e := "some command; ls -la"
	assert.Equal(t, result, e)

	run.PostCommands = []string{"some command 2"}
	cc = newConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result = strings.Join(cc.prepareCommand(), " ")
	e = "some command; ls -la; some command 2"
	assert.Equal(t, result, e)

	run = Run{
		PostCommands: []string{"some command 2"},
	}
	cc = newConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result = strings.Join(cc.prepareCommand(), " ")
	e = "ls -la; some command 2"
	assert.Equal(t, result, e)

	run.PostCommands = []string{"some command 2", "some command 3"}
	cc = newConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result = strings.Join(cc.prepareCommand(), " ")
	e = "ls -la; some command 2; some command 3"
	assert.Equal(t, result, e)
}

func TestBinBash(t *testing.T) {
	cc := newPureConfigConfigurator([]string{}, []string{})
	assert.Empty(t, cc.binBash())

	run := Run{
		PreCommands: []string{"some command"},
	}
	cc = newConfigConfigurator([]string{}, []string{}, &run)
	fmt.Println(cc.binBash())
	result := strings.Join(cc.binBash(), " ")
	e := "/bin/bash -c"
	assert.Equal(t, result, e)
}
