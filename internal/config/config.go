package config

import (
	"fmt"
	"os"
	"satellite/internal/entity"
	"satellite/internal/entity/docker"
	docker_compose "satellite/internal/entity/docker-compose"
	"sync"

	"satellite/internal/validator"

	"github.com/gookit/color"

	"satellite/pkg"

	"gopkg.in/yaml.v2"
)

var once sync.Once
var instance *Config

type Service struct {
	Name        string
	Description string
}

type Macros struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	List        []string `yaml:"commands"`
}

type DockerCompose struct {
	Run   []docker_compose.Run   `yaml:"run"`
	Up    []docker_compose.Up    `yaml:"up"`
	Down  []docker_compose.Down  `yaml:"down"`
	Exec  []docker_compose.Exec  `yaml:"exec"`
	Build []docker_compose.Build `yaml:"build"`
}

type Docker struct {
	Pulls []docker.Pull `yaml:"pull"`
	Runs  []docker.Run  `yaml:"run"`
	Execs []docker.Exec `yaml:"exec"`
}

type Config struct {
	Path     string
	Macros   []Macros      `yaml:"macros"`
	Services Docker        `yaml:"docker"`
	DCompose DockerCompose `yaml:"docker-compose"`
}

func (c *Config) GetMacros(name string) *Macros {
	for _, v := range c.Macros {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (c *Config) FindService(name string) entity.Runner {
	var service entity.Runner
	for _, v := range c.ServicesList() {
		if v.GetName() == name {
			service = v
			break
		}
	}

	valid := validator.NewValidator()
	errs, isValid := valid.Validate(service)
	if isValid {
		return service
	}
	for _, v := range errs {
		color.Red.Printf("Service %s error: %s\n", service.GetName(), v)
	}
	os.Exit(1)
	return nil
}

func (c *Config) ServicesList() map[string]entity.Runner {
	data := make(map[string]entity.Runner)

	for i := 0; i < len(c.Services.Pulls); i++ {
		item := &c.Services.Pulls[i]
		data[item.GetName()] = &c.Services.Pulls[i]
	}

	for i := 0; i < len(c.Services.Runs); i++ {
		item := &c.Services.Runs[i]
		data[item.GetName()] = &c.Services.Runs[i]
	}

	for i := 0; i < len(c.Services.Execs); i++ {
		item := &c.Services.Execs[i]
		data[item.GetName()] = &c.Services.Execs[i]
	}

	for i := 0; i < len(c.DCompose.Run); i++ {
		item := &c.DCompose.Run[i]
		data[item.GetName()] = &c.DCompose.Run[i]
	}

	for i := 0; i < len(c.DCompose.Up); i++ {
		item := &c.DCompose.Up[i]
		data[item.GetName()] = &c.DCompose.Up[i]
	}

	for i := 0; i < len(c.DCompose.Exec); i++ {
		item := &c.DCompose.Exec[i]
		data[item.GetName()] = &c.DCompose.Exec[i]
	}

	for i := 0; i < len(c.DCompose.Down); i++ {
		item := &c.DCompose.Down[i]
		data[item.GetName()] = &c.DCompose.Down[i]
	}

	for i := 0; i < len(c.DCompose.Build); i++ {
		item := &c.DCompose.Build[i]
		data[item.GetName()] = &c.DCompose.Build[i]
	}

	return data
}

func (c *Config) GetServices() []Service {
	var list []Service
	for _, v := range c.ServicesList() {
		list = append(list, Service{
			Name:        v.GetName(),
			Description: v.GetDescription(),
		})
	}
	return list
}

func mappingConfig(path string) *Config {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		color.Danger.Printf("error open config file, err: %s\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			color.Danger.Printf("error while closing file, err: %s\n", err)
			os.Exit(1)
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		color.Danger.Printf("error while decode config file, err: %s\n", err)
		os.Exit(1)
	}
	return config
}

func GetConfig() *Config {
	once.Do(func() {
		configName := os.Getenv("CONFIG_NAME")
		if configName == "" {
			configName = "satellite.yaml"
		}

		path := fmt.Sprintf("%s/%s", pkg.GetPwd(), configName)
		instance = mappingConfig(path)
	})

	return instance
}

func NewConfig(path string) *Config {
	return mappingConfig(path)
}
