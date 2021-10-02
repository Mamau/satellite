package docker

type docker struct {
	Name        string `yaml:"name" validate:"required,min=1"`
	Description string `yaml:"description"`
}

func (d *docker) GetName() string {
	return d.Name
}

func (d *docker) GetDescription() string {
	return d.Description
}
