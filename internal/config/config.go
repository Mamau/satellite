package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/mamau/satellite/internal/entity"

	"github.com/mamau/satellite/pkg"

	"gopkg.in/yaml.v2"
)

var once sync.Once
var instance *Config

type Service struct {
	Name        string
	Description string
}

type Macros struct {
	Name string   `yaml:"name"`
	List []string `yaml:"commands"`
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

func (c *Config) FindDCService(name string) *entity.DockerCompose {
	for _, v := range c.Services.ServiceDC {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (c *Config) FindExecService(name string) *entity.Exec {
	for _, v := range c.Services.ServiceExec {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (c *Config) FindPullService(name string) *entity.Pull {
	for _, v := range c.Services.ServicesPull {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (c *Config) FindRunService(name string) *entity.Run {
	for _, v := range c.Services.ServiceRun {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (c *Config) FindService(name string) entity.Runner {
	for _, v := range c.ServicesList() {
		if v.GetName() == name {
			return v
		}
	}
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
