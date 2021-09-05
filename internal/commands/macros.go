package commands

import (
	"strings"

	"github.com/gookit/color"
	"github.com/mamau/satellite/internal/config"
	"github.com/spf13/cobra"
)

var macrosCmd = &cobra.Command{
	Use:   "macros",
	Short: "Run group of commands",
	Long:  "Run group of commands",
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
					return
				}

				for _, v := range macros.List {
					cml := strings.Split(v, " ")
					if serviceName := c.FindService(cml[0]); serviceName != nil {
						serviceCmd.Run(cmd, cml)
						continue
					}
					color.Danger.Printf("Service %s not found\n", cml[0])
				}
			},
		})
	}
}
