package config

import (
	"strings"
	"testing"

	"github.com/mamau/starter/libs"
)

func TestGetMacros(t *testing.T) {
	c := getConfig("/testdata/starter")
	if m := c.GetMacros("test_not_found"); m != nil {
		t.Errorf("macros must be empty")
	}

	if m := c.GetMacros("test"); m == nil {
		t.Errorf("macros must be not empty")
	}
}

func TestGetService(t *testing.T) {
	c := getConfig("/testdata/starter")
	if s := c.GetService("php"); s.Name != "php" {
		t.Error("got wrong service")
	}

	if s := c.GetService("unknown"); s != nil {
		t.Error("unknown service must be nil")
	}
}

func TestGetServices(t *testing.T) {
	c := getConfig("/testdata/starter")
	if list := c.GetServices(); strings.Join(list, " ") != "php mysql" {
		t.Errorf("services expected %q", "php mysql")
	}
}

func TestGetClientConfig(t *testing.T) {
	fp := libs.GetPwd() + "/testdata/starter"
	result := GetClientConfig(fp)

	if result != fp+".yaml" {
		t.Errorf("file %s is not exist", fp)
	}

	fp = libs.GetPwd() + "/testdata/starter_not_exists"
	result = GetClientConfig(fp)
	if result != "" {
		t.Errorf("file %s not exists and return non empty string", fp)
	}
}

func getConfig(cn string) *Config {
	c := NewConfig(libs.GetPwd() + "/testdata/starter")
	c.Path = libs.GetPwd() + cn
	return GetConfig()
}
