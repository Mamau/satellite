package commands

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"satellite/internal/config"
	"satellite/internal/entity"
	docker_compose "satellite/internal/entity/docker-compose"
	"satellite/internal/informator"
)

var dockerComposeCmd = &cobra.Command{
	Use:     "docker-compose",
	Short:   "показать доступные опции докер композ секции",
	Long:    "позволяет показать какие опции можно использовать в конфиге докер композ секции",
	Example: "./sat docker-compose exec",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			showDCTypes()
			return
		}
		cmdName := args[0]
		ct := getDCCommand(cmdName)
		if ct == nil {
			showDCTypes()
			return
		}

		info := informator.NewInformator(ct)
		for i, v := range info.GetAll() {
			color.Cyan.Printf("%-25s%s\n", i, v)
		}
	},
}

func showDCTypes() {
	d := config.DockerCompose{}
	color.White.Println("Доступные команды для получения информации (попробуй ./sat docker pull):")
	for _, v := range d.GetTypes() {
		color.Cyan.Println(v)
	}
}

func getDCCommand(slug string) entity.Runner {
	switch slug {
	case "exec":
		return docker_compose.Exec{}
	case "run":
		return docker_compose.Run{}
	case "up":
		return docker_compose.Up{}
	case "build":
		return docker_compose.Build{}
	case "down":
		return docker_compose.Down{}
	default:
		return nil
	}
}
