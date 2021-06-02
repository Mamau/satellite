package config

type Macros struct {
	Name string   `yaml:"name"`
	List []string `yaml:"commands"`
}
