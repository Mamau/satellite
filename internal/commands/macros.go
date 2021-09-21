package commands

import (
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/mamau/satellite/internal/config"
	"github.com/spf13/cobra"
)

var macrosCmd = &cobra.Command{
	Use:   "macros",
	Short: "Run group of commands",
	Long:  "Run separate commands from service section in one macros command",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			color.Danger.Println("Error while running help of macros")
		}
	},
}

func init() {
	initMacrosSubCommand()
	rootCmd.AddCommand(macrosCmd)
}

func initMacrosSubCommand() {
	c := config.GetConfig()
	for _, item := range c.Macros {
		macrosCmd.AddCommand(&cobra.Command{
			Use:   item.Name,
			Short: item.Description,
			Long:  item.Description,
			Run: func(cmd *cobra.Command, args []string) {
				macrosName := cmd.Name()
				color.Cyan.Printf("Start %s\n", macrosName)

				macros := c.GetMacros(macrosName)

				if macros == nil {
					color.Danger.Println("Macros not found")
					os.Exit(1)
				}

				for _, v := range getServices(macros.List) {
					serviceCmd.Run(cmd, v)
				}
			},
		})
	}
}

func getServices(macrosList []string) [][]string {
	var commandList [][]string
	c := config.GetConfig()
	for _, v := range macrosList {
		cml := strings.Split(v, " ")
		if serviceName := c.FindService(cml[0]); serviceName != nil {
			commandList = append(commandList, cml)
			continue
		}
		color.Danger.Printf("Service %s not found\n", cml[0])
	}

	return commandList
}
