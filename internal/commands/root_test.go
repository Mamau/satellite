package commands

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestRunner struct {
	err error
}

func (s *TestRunner) Run() error {
	return s.err
}

func TestCheckDockerService(t *testing.T) {
	cases := []struct {
		Name           string
		CmdName        string
		ReplaceGateWay []string
		LookPath       func(file string) (string, error)
		Command        func(name string, arg ...string) CommandRunner
		HasError       bool
		Expected       func() (string, []string)
	}{
		{
			Name:           "сервис docker в системе есть",
			CmdName:        "docker",
			ReplaceGateWay: []string{"run"},
			LookPath: func(file string) (string, error) {
				return "", nil
			},
			Command: func(name string, arg ...string) CommandRunner {
				return nil
			},
			Expected: func() (string, []string) {
				return "docker", []string{"run"}
			},
			HasError: false,
		},
		{
			Name:           "сервис docker в системе отсутствует",
			CmdName:        "docker",
			ReplaceGateWay: []string{"run"},
			LookPath: func(file string) (string, error) {
				return "", errors.New("docker service not found")
			},
			Command: func(name string, arg ...string) CommandRunner {
				return nil
			},
			Expected: func() (string, []string) {
				return "", nil
			},
			HasError: true,
		},
		{
			Name:           "сервис docker-compose в системе отсутствует и нет compose 2nd версии",
			CmdName:        "docker-compose",
			ReplaceGateWay: []string{"run"},
			LookPath: func(file string) (string, error) {
				return "", errors.New("docker-compose service not found")
			},
			Command: func(name string, arg ...string) CommandRunner {
				testRunner := &TestRunner{
					err: errors.New("cant run"),
				}
				return testRunner
			},
			Expected: func() (string, []string) {
				return "", nil
			},
			HasError: true,
		},
		{
			Name:           "сервис docker-compose в системе отсутствует но есть вторая версия",
			CmdName:        "docker-compose",
			ReplaceGateWay: []string{"run"},
			LookPath: func(file string) (string, error) {
				return "", errors.New("docker-compose service not found")
			},
			Command: func(name string, arg ...string) CommandRunner {
				testRunner := &TestRunner{}
				return testRunner
			},
			Expected: func() (string, []string) {
				return "docker", []string{"compose", "run"}
			},
			HasError: false,
		},
	}
	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			runner := v.Command("docker", "compose")
			cmdName, args, err := checkDockerService(v.CmdName, v.ReplaceGateWay, v.LookPath, runner)
			if err == nil {
				assert.False(t, v.HasError)
			} else {
				assert.Error(t, err)
			}
			expectedCmdName, expectedArgs := v.Expected()
			assert.Equal(t, expectedCmdName, cmdName)
			assert.Equal(t, expectedArgs, args)
		})
	}
}
