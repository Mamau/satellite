package libs

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/mamau/starter/libs/services"

	"gopkg.in/yaml.v2"
)

var once sync.Once
var instance *Config

type Config struct {
	Services struct {
		services.Composer `yaml:"composer"`
		services.Yarn     `yaml:"yarn"`
		services.Gulp     `yaml:"gulp"`
		services.Bower    `yaml:"bower"`
	} `yaml:"services"`
}

func NewConfig() *Config {
	once.Do(func() {
		instance = &Config{}
	})

	return instance
}

func (c *Config) GetComposer() *services.Composer {
	return &c.Services.Composer
}

func (c *Config) GetYarn() *services.Yarn {
	return &c.Services.Yarn
}

func (c *Config) GetBower() *services.Bower {
	return &c.Services.Bower
}

func (c *Config) GetGulp() *services.Gulp {
	return &c.Services.Gulp
}

func GetConfig() *Config {
	fileName := getClientConfig()
	c := NewConfig()
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

func getClientConfig() string {
	for _, ext := range []string{"yaml", "yml"} {
		file := fmt.Sprintf("%s.%s", "starter", ext)
		if FileExists(file) {
			return file
		}
	}
	return ""
}
