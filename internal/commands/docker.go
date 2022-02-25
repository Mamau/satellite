package commands

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"satellite/internal/config"
	"satellite/internal/entity"
	"satellite/internal/entity/docker"
	"satellite/internal/informator"
)

var dockerCmd = &cobra.Command{
	Use:     "docker",
	Short:   "показать доступные опции докер секции",
	Long:    "позволяет показать какие опции можно использовать в конфиге докер секции",
	Example: "./sat docker pull",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			showTypes()
			return
		}
		cmdName := args[0]
		ct := getCommand(cmdName)
		if ct == nil {
			showTypes()
			return
		}

		info := informator.NewInformator(ct)
		for i, v := range info.GetAll() {
			color.Cyan.Printf("%-25s%s\n", i, v)
		}
	},
}

func showTypes() {
	d := config.Docker{}
	color.White.Println("Доступные команды для получения информации (попробуй ./sat docker pull):")
	for _, v := range d.GetTypes() {
		color.Cyan.Println(v)
	}
}

func getCommand(slug string) entity.Runner {
	switch slug {
	case "exec":
		return docker.Exec{}
	case "run":
		return docker.Run{}
	case "pull":
		return docker.Pull{}
	default:
		return nil
	}
}
