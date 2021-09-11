package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mamau/satellite/pkg"
)

func TestGetMacros(t *testing.T) {
	c := getConfig("/testdata/satellite")
	assert.Empty(t, c.GetMacros("test_not_found"))

	assert.NotEmpty(t, c.GetMacros("test"))
}

func TestGetServices(t *testing.T) {
	c := getConfig("/testdata/satellite")
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

func TestGetClientConfig(t *testing.T) {
	fp := pkg.GetPwd() + "/testdata/satellite"
	result := GetClientConfig(fp)

	assert.Equal(t, fp+".yaml", result)

	fp = pkg.GetPwd() + "/testdata/satellite_not_exists"
	result = GetClientConfig(fp)
	assert.Empty(t, result)
}

func getConfig(cn string) *Config {
	c := NewConfig(pkg.GetPwd() + "/testdata/satellite")
	c.Path = pkg.GetPwd() + cn
	return GetConfig()
}
