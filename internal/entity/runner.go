package entity

type ExecCommand string

const (
	DOCKER           ExecCommand = "docker"
	DOCKER_COMPOSE   ExecCommand = "docker-compose"
	DOCKER_COMPOSE_2 ExecCommand = "docker compose"
)

type Runner interface {
	GetExecCommand() string
	ToCommand(args []string) []string
	GetName() string
	GetDescription() string
	//GetParams() []string TODO: make a list of available commands docker
}
