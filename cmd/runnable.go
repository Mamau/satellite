package cmd

type Runnable interface {
	CollectCommand() []string
	GetBeginCommand() []string
}
