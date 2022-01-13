package docker_compose

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDCToCommand(t *testing.T) {
	cases := []struct {
		Name      string
		Expected  string
		Path      string
		Command   string
		Verbose   bool
		MultiPath []string
	}{
		{
			Name:      "путей нет",
			Expected:  "",
			Command:   "",
			Path:      "",
			Verbose:   false,
			MultiPath: nil,
		},
		{
			Name:      "добавлено один путь",
			Expected:  "--file /some/path --verbose up",
			Path:      "/some/path",
			Command:   "up",
			Verbose:   true,
			MultiPath: nil,
		},
		{
			Name:     "добавлено несколько путей",
			Expected: "--file /path/1 --file /path/2 --verbose up",
			Path:     "",
			Command:  "up",
			Verbose:  true,
			MultiPath: []string{
				"/path/1",
				"/path/2",
			},
		},
		{
			Name:     "добавлено несколько путей и дефолтный",
			Expected: "--file /some/path --file /path/1 --file /path/2 --verbose up",
			Path:     "/some/path",
			Verbose:  true,
			Command:  "up",
			MultiPath: []string{
				"/path/1",
				"/path/2",
			},
		},
	}
	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			dc := DockerCompose{
				MultiPath: v.MultiPath,
				Path:      v.Path,
				Verbose:   v.Verbose,
			}
			result := strings.Join(dc.ToCommand([]string{v.Command}), " ")
			assert.Equal(t, v.Expected, result)
		})
	}
}

func TestGetProjectDirectory(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetProjectDirectory())

	dc.ProjectDirectory = "/some/dir"
	assert.Equal(t, dc.GetProjectDirectory(), "--project-directory /some/dir")
}

func TestGetPath(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetPath())

	dc.Path = "/some/path"
	assert.Equal(t, dc.GetPath(), "--file /some/path")
}

func TestGetMultiPath(t *testing.T) {
	cases := []struct {
		Name      string
		Expected  string
		MultiPath []string
	}{
		{
			Name:      "путей нет",
			Expected:  "",
			MultiPath: nil,
		},
		{
			Name:     "добавлено несколько путей",
			Expected: "--file /path/1 --file /path/2",
			MultiPath: []string{
				"/path/1",
				"/path/2",
			},
		},
	}

	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			dc := DockerCompose{
				MultiPath: v.MultiPath,
			}

			assert.Equal(t, v.Expected, dc.GetMultiPath())
		})
	}
}

func TestGetProjectName(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetProjectName())

	dc.ProjectName = "some-p-name"
	assert.Equal(t, dc.GetProjectName(), "--project-name some-p-name")
}

func TestGetLogLevel(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetLogLevel())

	dc.LogLevel = "DEBUG"
	assert.Equal(t, dc.GetLogLevel(), "--log-level DEBUG")
}

func TestGetVerbose(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetVerbose())

	dc.Verbose = true
	assert.Equal(t, dc.GetVerbose(), "--verbose")
}

func TestDCGetName(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetName())

	dc.Name = "some-name"
	assert.Equal(t, dc.GetName(), "some-name")
}

func TestDCGetDescription(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetDescription())

	dc.Description = "some description"
	assert.Equal(t, dc.GetDescription(), "some description")
}

func TestDCGetExecCommand(t *testing.T) {
	dc := DockerCompose{}
	assert.Equal(t, dc.GetExecCommand(), "docker-compose")
}
