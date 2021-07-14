package docker

type Exec string

const (
	DOCKER         Exec = "docker"
	DOCKER_COMPOSE Exec = "docker-compose"
	PULL           Exec = "pull"
	DAEMON         Exec = "daemon"
	RUN            Exec = "run"
)

var List = []Exec{
	DOCKER,
	DOCKER_COMPOSE,
	PULL,
	DAEMON,
	RUN,
}

func (e Exec) String() string {
	return string(e)
}

func (e Exec) IsAllowed() bool {
	for _, n := range List {
		if e == n {
			return true
		}
	}
	return false
}
