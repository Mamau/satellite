package commands

type Runnable interface {
	CollectCommand() []string
}
