package strategy

type ExecCommand string

const (
	DOCKER         ExecCommand = "docker"
	DOCKER_COMPOSE ExecCommand = "docker-compose"
)
