package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/mamau/starter/config/yarn"

	"github.com/mamau/starter/config/composer"
	"github.com/mamau/starter/libs"

	"gopkg.in/yaml.v2"
)

var once sync.Once
var instance *Config

type Config struct {
	Path     string
	Macros   []string `yaml:"macros"`
	Commands struct {
		*composer.Composer `yaml:"composer"`
		*yarn.Yarn         `yaml:"yarn"`
		*Bower             `yaml:"bower"`
	} `yaml:"commands"`
}

func NewConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Path: libs.GetPwd() + "/starter",
		}
	})

	return instance
}

func (c *Config) GetMacrosGroup() []string {
	return c.Macros
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
