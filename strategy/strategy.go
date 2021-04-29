package strategy

type Strategy interface {
	ToCommand() []string
}
