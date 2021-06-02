package config

import (
	"strings"
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

func TestGetService(t *testing.T) {
	c := getConfig("/testdata/satellite")
	if s := c.GetService("php"); s.Name != "php" {
		t.Error("got wrong service")
	}

	if s := c.GetService("unknown"); s != nil {
		t.Error("unknown service must be nil")
	}
}

func TestGetServices(t *testing.T) {
	c := getConfig("/testdata/satellite")
	if list := c.GetServices(); strings.Join(list, " ") != "php mysql" {
		t.Errorf("services expected %q", "php mysql")
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
