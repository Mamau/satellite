package commands

import (
	"github.com/gookit/color"
	"github.com/mamau/satellite/internal/config"
	"github.com/mamau/satellite/pkg"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		var serviceName string
		var arguments []string

		if len(args) < 1 {
			color.Red.Printf("You must pass a service name\n")
			return
		}

		serviceName = args[0]

		if len(args) >= 2 {
			arguments = args[1:]
		}

		color.Cyan.Printf("Start %s\n", serviceName)

		s := config.GetConfig().FindService(serviceName)

		pkg.RunCommandAtPTY(Docker(s, arguments))
	},
}
