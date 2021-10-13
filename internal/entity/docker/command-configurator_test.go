package docker

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientCommand(t *testing.T) {
	run := Run{}
	cc := NewConfigConfigurator([]string{}, []string{}, &run)
	result := strings.Join(cc.GetClientCommand(), " ")
	assert.Empty(t, result)

	cc = NewConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result = strings.Join(cc.GetClientCommand(), " ")
	assert.Equal(t, result, "ls -la")

	run.PreCommands = []string{"some command"}
	cc = NewConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result = strings.Join(cc.GetClientCommand(), " ")
	assert.Equal(t, result, "/bin/bash -c some command; ls -la")
}

func TestPrepareCommand(t *testing.T) {
	run := Run{}
	cc := NewConfigConfigurator([]string{}, []string{}, &run)
	assert.Empty(t, cc.prepareCommand())

	run.PreCommands = []string{"some command"}
	cc = NewConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result := strings.Join(cc.prepareCommand(), " ")
	e := "some command; ls -la"
	assert.Equal(t, result, e)

	run.PostCommands = []string{"some command 2"}
	cc = NewConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result = strings.Join(cc.prepareCommand(), " ")
	e = "some command; ls -la; some command 2"
	assert.Equal(t, result, e)

	run = Run{
		PostCommands: []string{"some command 2"},
	}
	cc = NewConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result = strings.Join(cc.prepareCommand(), " ")
	e = "ls -la; some command 2"
	assert.Equal(t, result, e)

	run.PostCommands = []string{"some command 2", "some command 3"}
	cc = NewConfigConfigurator([]string{}, []string{"ls -la"}, &run)
	result = strings.Join(cc.prepareCommand(), " ")
	e = "ls -la; some command 2; some command 3"
	assert.Equal(t, result, e)
}

func TestBinBash(t *testing.T) {
	run := Run{}
	cc := NewConfigConfigurator([]string{}, []string{}, &run)
	assert.Empty(t, cc.binBash())

	run.PreCommands = []string{"some command"}
	cc = NewConfigConfigurator([]string{}, []string{}, &run)
	fmt.Println(cc.binBash())
	result := strings.Join(cc.binBash(), " ")
	e := "/bin/bash -c"
	assert.Equal(t, result, e)
}
