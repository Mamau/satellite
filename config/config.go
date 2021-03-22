package config

import (
	"fmt"
	"github.com/mamau/starter/libs"
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

var once sync.Once
var instance *Config

type Config struct {
	Path     string
	Services struct {
		*Composer `yaml:"composer"`
		*Yarn     `yaml:"yarn"`
		*Bower    `yaml:"bower"`
	} `yaml:"services"`
}

func NewConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Path: libs.GetPwd() + "/starter",
		}
	})

	return instance
}

func (c *Config) GetComposer() *Composer {
	return c.Services.Composer
}

func (c *Config) GetYarn() *Yarn {
	return c.Services.Yarn
}

func (c *Config) GetBower() *Bower {
	return c.Services.Bower
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
