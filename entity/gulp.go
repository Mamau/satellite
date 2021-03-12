package entity

type Gulp struct {
	Command
}

func (g *Gulp) CollectCommand() []string {
	var fullCommand []string
	commandParts := [][]string{
		g.workDirVolume(),
		g.projectVolume(),
		{g.getImage()},
		{g.fullCommand()},
	}
	for _, command := range commandParts {
		fullCommand = append(fullCommand, command...)
	}

	return fullCommand
}
