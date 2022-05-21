package commands

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	docker_compose "satellite/internal/entity/docker-compose"
	"satellite/internal/mock"
	"satellite/internal/usecase/collector"
	"testing"
)

func TestCheckCommand(t *testing.T) {
	cases := []struct {
		Name      string
		CmdArgs   []string
		Collector func() collector.Finder
		Expected  error
	}{
		{
			Name:    "команда существую и ошибки нет",
			CmdArgs: []string{"testName"},
			Collector: func() collector.Finder {
				runner := docker_compose.Run{}
				runner.Name = "testName"
				ctrl := gomock.NewController(GinkgoT())
				finder := mock.NewMockFinder(ctrl)
				finder.EXPECT().FindCommand("testName").Return(runner).AnyTimes()
				return finder
			},
			Expected: nil,
		},
		{
			Name:    "команда не существую и ошибка есть",
			CmdArgs: []string{"testNoName"},
			Collector: func() collector.Finder {
				ctrl := gomock.NewController(GinkgoT())
				finder := mock.NewMockFinder(ctrl)
				finder.EXPECT().FindCommand("testNoName").Return(nil).AnyTimes()
				return finder
			},
			Expected: errors.New("команда testNoName не найдена в разделе макросов"),
		},
		{
			Name:    "команда не передана и ошибки нет",
			CmdArgs: []string{},
			Collector: func() collector.Finder {
				return nil
			},
			Expected: nil,
		},
	}
	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			result := checkCommand(v.Collector(), v.CmdArgs)
			assert.Equal(t, v.Expected, result)
		})
	}
}

func TestPrepareArgs(t *testing.T) {
	args := []string{"composer-cmd-name", "composer", "i", "--ignore-platform-reqs"}
	sn, command := prepareArgs(args)
	assert.Equal(t, sn, args[0])
	assert.Equal(t, command, args[1:])

	args = []string{"composer-cmd-name"}
	sn, command = prepareArgs(args)
	assert.Equal(t, sn, args[0])
	assert.Empty(t, command)
}
