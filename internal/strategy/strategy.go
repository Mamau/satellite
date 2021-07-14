package strategy

type Strategy interface {
	ToCommand() []string
	GetExecCommand() string
	GetContext() CommandContext
}
