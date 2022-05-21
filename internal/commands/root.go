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

var serviceCollector *collector.Service

func InitCommands() {
	serviceCollector = collector.NewService(config.GetConfig())

	registerCommands()
	initServiceCommand()
	initMacrosSubCommand()

	execute()
}

func registerCommands() {
	rootCmd.AddCommand(macrosCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(dockerCmd)
	rootCmd.AddCommand(dockerComposeCmd)
}

func initServiceCommand() {
	for _, service := range serviceCollector.ServicesList() {
		rootCmd.AddCommand(&cobra.Command{
			Use:                service.GetName(),
			Short:              service.GetDescription(),
			Long:               service.GetDescription(),
			DisableFlagParsing: true,
			PreRun: func(cmd *cobra.Command, args []string) {
				color.Cyan.Printf("Start %s\n", cmd.Name())
			},
			Run: func(cmd *cobra.Command, args []string) {
				s := serviceCollector.FindCommand(cmd.Name())

				pkg.RunCommandAtPTY(serviceCollector.ExecuteCommand(s, args))
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
