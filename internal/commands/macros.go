package commands

import (
	"strings"

	"github.com/mamau/satellite/internal/config"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var macrosCmd = &cobra.Command{
	Use:   "macros",
	Short: "Run group of commands",
	Long:  "Run group of commands",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Danger.Println("You should pass macros name")
			return
		}

		macrosName := args[0]

		var cl []string

		c := config.GetConfig()
		macros := c.GetMacros(macrosName)

		if macros == nil {
			color.Danger.Println("Macros not found")
			return
		}

		for _, v := range macros.List {
			cml := strings.Split(v, " ")
			if serviceName := c.GetService(cml[0]); serviceName != nil {
				cl = append(cl, cml[0])
				serviceCmd.Run(cmd, cml)
			}
		}

		if len(cl) == 0 {
			color.Danger.Println("Commands in group not found")
		}
	},
}

func init() {
	rootCmd.AddCommand(macrosCmd)
}
