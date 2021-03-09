package composer

type Config struct {
	Repository `yaml:"repository"`
}

func (c *Config) GetRepository() *Repository {
	return &c.Repository
}
