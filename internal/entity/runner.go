package entity

type ExecCommand string

const (
	DOCKER         ExecCommand = "docker"
	DOCKER_COMPOSE ExecCommand = "docker-compose"
)

type Runner interface {
	GetExecCommand() string
	ToCommand(args []string) []string
	GetName() string
	GetDescription() string
}
