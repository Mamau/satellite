package entity

type Bower struct {
	Command
}

func (b *Bower) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		b.workDirVolume(),
		b.projectVolume(),
		{b.getImage()},
		{b.fullCommand()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}
