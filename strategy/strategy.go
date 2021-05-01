package strategy

import "context"

const PullType = "pull"
const DaemonType = "daemon"
const RunType = "run"

type Strategy interface {
	ToCommand() []string
	GetContext() context.Context
}
