package commands

import (
	"satellite/internal/entity"
	"satellite/internal/entity/docker"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetServices(t *testing.T) {
	var emptyMacrosList []string
	r := getServices(finder, emptyMacrosList)
	assert.Empty(t, r)

	macrosList := []string{"yarn install", "php -v"}
	r = getServices(finder, macrosList)
	assert.NotEmpty(t, r)
	assert.Len(t, r, 2)
	assert.Equal(t, r[1][0], "php")
	assert.Equal(t, r[1][1], "-v")
}

func finder(name string) entity.Runner {
	e := &docker.Run{Name: "yarn"}
	return e
}
