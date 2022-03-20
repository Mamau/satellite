package commands

import (
	"fmt"
	"os"
	"satellite/internal/usecase/collector"
	"satellite/pkg"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var macrosCmd = &cobra.Command{
	Use:     "macros",
	Short:   "Запускает группу команд",
	Long:    "Запускает группу команд которые описаны в docker-compose или docker секции",
	Example: "./sat macros init",
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkCommand(serviceCollector, args); err != nil {
			color.Danger.Println(err.Error())
			os.Exit(1)
		}

		if err := cmd.Help(); err != nil {
			color.Danger.Println("Ошибка вызова команды помощи")
		}
	},
}

func initMacrosSubCommand() {
	for _, item := range serviceCollector.GetMacrosList() {
		macrosCmd.AddCommand(&cobra.Command{
			Use:   item.Name,
			Short: item.Description,
			Long:  item.Description,
			PreRun: func(cmd *cobra.Command, args []string) {
				color.Cyan.Printf("Start macros %q\n", cmd.Name())
			},
			Run: func(cmd *cobra.Command, args []string) {
				for _, v := range serviceCollector.GetMacrosCommands(item.List) {
					serviceName, arguments := prepareArgs(v)
					s := serviceCollector.FindCommand(serviceName)

					pkg.RunCommandAtPTY(serviceCollector.ExecuteCommand(s, arguments))
				}
			},
		})
	}
}

func checkCommand(finder collector.Finder, args []string) error {
	if len(args) > 0 {
		serviceName := args[0]

		if s := finder.FindCommand(serviceName); s == nil {
			return fmt.Errorf("команда %s не найдена в разделе макросов", serviceName)
		}
	}

	return nil
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
