package strategy

type Strategy interface {
	ToCommand() []string
	GetContext() CommandContext
}
