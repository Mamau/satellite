package config

import (
	"testing"

	"github.com/mamau/satellite/pkg"
)

func TestGetMacros(t *testing.T) {
	c := getConfig("/testdata/satellite")
	if m := c.GetMacros("test_not_found"); m != nil {
		t.Errorf("macros must be empty")
	}

	if m := c.GetMacros("test"); m == nil {
		t.Errorf("macros must be not empty")
	}
}

func TestGetServices(t *testing.T) {
	c := getConfig("/testdata/satellite")
	services := c.GetServices()
	e := "php"
	if services[0].Name != e {
		t.Errorf("services #1 expected name %q", e)
	}

	ed := "some description"
	if services[0].Description != ed {
		t.Errorf("services #1 expected description %q", ed)
	}

	e2 := "mysql"
	if services[1].Name != e2 {
		t.Errorf("services #2 expected name %q", e2)
	}

	ed2 := "some mysql description"
	if services[1].Description != ed2 {
		t.Errorf("services #2 expected description %q", ed2)
	}
}

func TestGetClientConfig(t *testing.T) {
	fp := pkg.GetPwd() + "/testdata/satellite"
	result := GetClientConfig(fp)

	if result != fp+".yaml" {
		t.Errorf("file %s is not exist", fp)
	}

	fp = pkg.GetPwd() + "/testdata/satellite_not_exists"
	result = GetClientConfig(fp)
	if result != "" {
		t.Errorf("file %s not exists and return non empty string", fp)
	}
}

func getConfig(cn string) *Config {
	c := NewConfig(pkg.GetPwd() + "/testdata/satellite")
	c.Path = pkg.GetPwd() + cn
	return GetConfig()
}
