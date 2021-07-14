package docker

type LogLevel string

const (
	DEBUG    LogLevel = "DEBUG"
	WARNING  LogLevel = "WARNING"
	ERROR    LogLevel = "ERROR"
	CRITICAL LogLevel = "CRITICAL"
)

var LogLevelList = []LogLevel{
	DEBUG,
	WARNING,
	ERROR,
	CRITICAL,
}

func (l LogLevel) String() string {
	return string(l)
}

func (l LogLevel) IsAllowed() bool {
	for _, n := range LogLevelList {
		if l == n {
			return true
		}
	}
	return false
}
