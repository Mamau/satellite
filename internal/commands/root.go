package commands

import (
	"fmt"
	"os"
	"satellite/internal/usecase/collector"
	"satellite/pkg"

	"satellite/internal/config"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sat",
	Short: "All command",
	Long:  "Show all command",
}

func InitCommands() {
	registerCommands()
	initServiceCommand()
	initMacrosSubCommand()

	execute()
}

func registerCommands() {
	rootCmd.AddCommand(macrosCmd)
	rootCmd.AddCommand(updateCmd)
}

func initServiceCommand() {
	serv := collector.NewService(config.GetConfig())
	for _, service := range serv.ServicesList() {
		rootCmd.AddCommand(&cobra.Command{
			Use:                service.GetName(),
			Short:              service.GetDescription(),
			Long:               service.GetDescription(),
			DisableFlagParsing: true,
			PreRun: func(cmd *cobra.Command, args []string) {
				color.Cyan.Printf("Start %s\n", cmd.Name())
			},
			Run: func(cmd *cobra.Command, args []string) {
				s := serv.FindCommand(cmd.Name())

				pkg.RunCommandAtPTY(serv.ExecuteCommand(s, args))
			},
		})
	}
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
