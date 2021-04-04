package cmd

import (
	"fmt"

	"github.com/mamau/starter/entity"
	"github.com/mamau/starter/libs"
	"github.com/mamau/starter/services"

	"github.com/mamau/starter/config"

	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("You should pass args")
			return
		}

		serviceName := args[0]

		config := config.GetConfig()
		s := config.GetService(serviceName)

		ser := entity.NewService(s, args[1:])
		collector := services.NewCollector(ser)
		libs.RunCommandAtPTY(Docker(collector))
	},
}
