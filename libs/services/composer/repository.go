package composer

import "strings"

type Repository struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
}

func (r *Repository) ToCommand() string {
	if r.Url == "" || r.Username == "" || r.Token == "" {
		return ""
	}

	args := []string{
		"composer",
		"config",
		r.Url,
		r.Username,
		r.Token,
	}
	return strings.Join(args, " ")
}
