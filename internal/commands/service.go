package commands

import (
	"os"

	"satellite/internal/config"
	"satellite/pkg"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		serviceName, arguments := prepareArgs(args)

		color.Cyan.Printf("Start %q\n", serviceName)

		s := config.GetConfig().FindService(serviceName)

		pkg.RunCommandAtPTY(Docker(s, arguments))
	},
}

func prepareArgs(args []string) (string, []string) {
	var serviceName string
	var arguments []string

	if len(args) < 1 {
		color.Red.Printf("You must pass a service name\n")
		os.Exit(1)
	}

	serviceName = args[0]

	if len(args) >= 2 {
		arguments = args[1:]
	}

	return serviceName, arguments
}
