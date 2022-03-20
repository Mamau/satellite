package config

type Macros struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	List        []string `yaml:"commands"`
}
