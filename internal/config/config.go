package config

import (
	"fmt"
	"os"
	"sync"

	"satellite/internal/validator"

	"github.com/gookit/color"

	"satellite/internal/entity"

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

type Services struct {
	ServicesPull []entity.Pull          `yaml:"pull"`
	ServiceRun   []entity.Run           `yaml:"run"`
	ServiceExec  []entity.Exec          `yaml:"exec"`
	ServiceDC    []entity.DockerCompose `yaml:"docker-compose"`
}

type Config struct {
	Path     string
	Macros   []Macros `yaml:"macros"`
	Services Services `yaml:"services"`
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

	for i := 0; i < len(c.Services.ServicesPull); i++ {
		item := &c.Services.ServicesPull[i]
		data[item.GetName()] = &c.Services.ServicesPull[i]
	}

	for i := 0; i < len(c.Services.ServiceRun); i++ {
		item := &c.Services.ServiceRun[i]
		data[item.GetName()] = &c.Services.ServiceRun[i]
	}

	for i := 0; i < len(c.Services.ServiceExec); i++ {
		item := &c.Services.ServiceExec[i]
		data[item.GetName()] = &c.Services.ServiceExec[i]
	}

	for i := 0; i < len(c.Services.ServiceDC); i++ {
		item := &c.Services.ServiceDC[i]
		data[item.GetName()] = &c.Services.ServiceDC[i]
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
