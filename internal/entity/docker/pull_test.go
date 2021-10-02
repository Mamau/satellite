package docker

import (
	"satellite/internal/entity"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPullGetExecCommand(t *testing.T) {
	pull := Pull{}
	assert.Equal(t, pull.GetExecCommand(), string(entity.DOCKER))
}

func TestPullGetDescription(t *testing.T) {
	pull := Pull{}
	assert.Empty(t, pull.GetDescription())

	pull.Description = "some description"
	assert.Equal(t, pull.GetDescription(), "some description")
}

func TestGetQuiet(t *testing.T) {
	pull := Pull{}
	assert.Empty(t, pull.GetQuiet())

	pull.Quiet = true
	assert.Equal(t, pull.GetQuiet(), "--quiet")
}

func TestGetDisableContentTrust(t *testing.T) {
	pull := Pull{}
	assert.Empty(t, pull.GetDisableContentTrust())

	pull.DisableContentTrust = "false"
	assert.Equal(t, pull.GetDisableContentTrust(), "--disable-content-trust=false")
}

func TestPullGetImage(t *testing.T) {
	pull := Pull{}
	assert.Empty(t, pull.GetImage())

	pull.Image = "some-image"
	assert.Equal(t, pull.GetImage(), "some-image")

	pull.Version = "3"
	assert.Equal(t, pull.GetImage(), "some-image:3")
}

func TestGetAllTags(t *testing.T) {
	pull := Pull{}
	assert.Empty(t, pull.GetAllTags())

	pull.AllTags = true
	assert.Equal(t, pull.GetAllTags(), "--all-tags")
}

func TestPullGetName(t *testing.T) {
	pull := Pull{}
	assert.Empty(t, pull.GetName())

	pull.Name = "img-name"
	assert.Equal(t, pull.GetName(), "img-name")
}

func TestPullToCommand(t *testing.T) {
	pull := Pull{}
	result := strings.Join(pull.ToCommand([]string{}), " ")
	assert.Equal(t, result, "pull")

	pull.DisableContentTrust = "false"
	pull.Image = "my-img"
	pull.Version = "1.4"
	pull.AllTags = true
	result = strings.Join(pull.ToCommand([]string{"pwd"}), " ")
	assert.Equal(t, result, "pull --disable-content-trust=false --all-tags my-img:1.4")
}
