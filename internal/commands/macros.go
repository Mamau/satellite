package commands

import (
	"os"
	"satellite/internal/config"
	"satellite/internal/usecase/collector"
	"satellite/pkg"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var macrosCmd = &cobra.Command{
	Use:     "macros",
	Short:   "Run group of commands",
	Long:    "Run separate commands from common section in one macros command",
	Example: "./sat macros init",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			color.Danger.Println("Error while running help of macros")
		}
	},
}

func initMacrosSubCommand() {
	service := collector.NewService(config.GetConfig())
	for _, item := range service.GetMacrosList() {
		macrosCmd.AddCommand(&cobra.Command{
			Use:   item.Name,
			Short: item.Description,
			Long:  item.Description,
			PreRun: func(cmd *cobra.Command, args []string) {
				color.Cyan.Printf("Start macros %q\n", cmd.Name())
			},
			Run: func(cmd *cobra.Command, args []string) {
				for _, v := range service.GetMacrosCommands(item.List) {
					serviceName, arguments := prepareArgs(v)
					serv := collector.NewService(config.GetConfig())
					s := serv.FindCommand(serviceName)

					pkg.RunCommandAtPTY(serv.ExecuteCommand(s, arguments))
				}
			},
		})
	}
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
