package commands

import (
	"testing"

	"satellite/internal/entity"

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
	e := &entity.Run{Name: "yarn"}
	return e
}
