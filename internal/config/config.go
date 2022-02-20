package config

import (
	"fmt"
	"github.com/gookit/color"
	"gopkg.in/yaml.v2"
	"os"
	"satellite/pkg"
)

var config *Config

type Config struct {
	Macros   []Macros      `yaml:"macros"`
	Services Docker        `yaml:"docker"`
	DCompose DockerCompose `yaml:"docker-compose"`
}

func (c *Config) GetMacros() []Macros {
	return c.Macros
}

func (c *Config) GetDocker() Commander {
	return &c.Services
}

func (c *Config) GetDockerCompose() Commander {
	return &c.DCompose
}

func GetConfig() *Config {
	if config == nil {
		configName := getEnv("CONFIG_NAME", "satellite.yaml")

		path := fmt.Sprintf("%s/%s", pkg.GetPwd(), configName)

		file, err := os.Open(path)
		if err != nil {
			color.Danger.Printf("error open config file, err: %s\n", err)
			return nil
		}

		defer func() {
			if err := file.Close(); err != nil {
				color.Danger.Printf("error while closing file, err: %s\n", err)
				os.Exit(1)
			}
		}()

		if err = yaml.NewDecoder(file).Decode(&config); err != nil {
			color.Danger.Printf("error while decode config file, err: %s\n", err)
			return nil
		}
	}

	return config
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
