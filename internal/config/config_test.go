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

	if services[0].Name == "php" {
		ed := "some description"
		assert.Equal(t, ed, services[0].Description)
		ed2 := "some mysql description"
		assert.Equal(t, ed2, services[1].Description)
	} else {
		ed := "some mysql description"
		assert.Equal(t, ed, services[0].Description)
		ed2 := "some description"
		assert.Equal(t, ed2, services[1].Description)
	}
}

func getConfig() *Config {
	return NewConfig(pkg.GetPwd() + "/testdata/satellite.yaml")
}
