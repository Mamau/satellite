package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/mamau/starter/config/docker"

	"github.com/mamau/starter/config/yarn"

	"github.com/mamau/starter/config/composer"
	"github.com/mamau/starter/libs"

	"gopkg.in/yaml.v2"
)

var once sync.Once
var instance *Config

type Config struct {
	Path     string
	Macros   []Macros        `yaml:"macros"`
	Services []docker.Docker `yaml:"services"`
	Commands struct {
		*composer.Composer `yaml:"composer"`
		*yarn.Yarn         `yaml:"yarn"`
		*Bower             `yaml:"bower"`
	} `yaml:"commands"`
}

func (c *Config) GetMacros(name string) *Macros {
	for _, v := range c.Macros {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (c *Config) GetService(name string) *docker.Docker {
	for _, v := range c.Services {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (c *Config) GetServices() []string {
	var list []string
	for _, v := range c.Services {
		list = append(list, v.Name)
	}
	return list
}

func NewConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Path: libs.GetPwd() + "/starter",
		}
	})

	return instance
}

func (c *Config) GetComposer() *composer.Composer {
	return c.Commands.Composer
}

func (c *Config) GetYarn() *yarn.Yarn {
	return c.Commands.Yarn
}

func (c *Config) GetBower() *Bower {
	return c.Commands.Bower
}

func GetConfig() *Config {
	c := NewConfig()
	fileName := GetClientConfig(c.Path)
	if fileName == "" {
		return c
	}

	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	if err := yaml.Unmarshal(buf, c); err != nil {
		log.Fatalln(err)
	}

	return c
}

func GetClientConfig(filePath string) string {
	for _, ext := range []string{"yaml", "yml"} {
		file := fmt.Sprintf("%s.%s", filePath, ext)
		if libs.FileExists(file) {
			return file
		}
	}
	return ""
}
