package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"satellite/pkg"
)

func TestGetMacros(t *testing.T) {
	c := getConfig()
	assert.Empty(t, c.GetMacros("test_not_found"))

	assert.NotEmpty(t, c.GetMacros("test"))
}

func TestGetServices(t *testing.T) {
	c := getConfig()
	services := c.GetServices()
	e := "php"
	assert.Equal(t, e, services[0].Name)

	ed := "some description"
	assert.Equal(t, ed, services[0].Description)

	e2 := "mysql"
	assert.Equal(t, e2, services[1].Name)

	ed2 := "some mysql description"
	assert.Equal(t, ed2, services[1].Description)
}

func getConfig() *Config {
	return NewConfig(pkg.GetPwd() + "/testdata/satellite.yaml")
}
