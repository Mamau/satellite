package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/mamau/satellite/pkg"

	"github.com/mamau/satellite/internal/config/docker"

	"gopkg.in/yaml.v2"
)

var once sync.Once
var instance *Config

type Macros struct {
	Name string   `yaml:"name"`
	List []string `yaml:"commands"`
}

type Config struct {
	Path     string
	Macros   []Macros        `yaml:"macros"`
	Services []docker.Docker `yaml:"services"`
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

func NewConfig(path string) *Config {
	once.Do(func() {
		instance = &Config{
			Path: path,
		}
	})

	return instance
}

func GetConfig() *Config {
	path := fmt.Sprintf("%s/satellite", pkg.GetPwd())
	c := NewConfig(path)
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
		if pkg.FileExists(file) {
			return file
		}
	}
	return ""
}
