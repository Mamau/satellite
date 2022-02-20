package commands

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"satellite/internal/entity/docker"
	"satellite/internal/informator"
)

var dockerCmd = &cobra.Command{
	Use:     "docker",
	Short:   "показать доступные опции докер секции exec/",
	Long:    "позволяет показать какие опции можно использовать в конфиге докер секции",
	Example: "./sat docker run",
	Run: func(cmd *cobra.Command, args []string) {
		exec := docker.Exec{}
		info := informator.NewInformator(exec)
		for i, v := range info.GetAll() {
			color.Cyan.Printf("%-20s%s\n", i, v)
		}
	},
}
