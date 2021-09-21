package commands

import (
	"testing"

	"github.com/mamau/satellite/internal/config"
	"github.com/mamau/satellite/pkg"

	"github.com/stretchr/testify/assert"
)

func TestGetServices(t *testing.T) {
	var emptyMacrosList []string
	getConfig("/testdata/satellite")
	r := getServices(emptyMacrosList)
	assert.Empty(t, r)

	macrosList := []string{"yarn install", "php -v"}
	r = getServices(macrosList)
	assert.NotEmpty(t, r)
	assert.Len(t, r, 2)
	assert.Equal(t, r[1][0], "php")
	assert.Equal(t, r[1][1], "-v")
}

func getConfig(cn string) *config.Config {
	c := config.NewConfig(pkg.GetPwd() + "/testdata/satellite")
	c.Path = pkg.GetPwd() + cn
	return config.GetConfig()
}
