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
		var macrosName string

		if len(args) < 1 {
			color.Red.Printf("You should pass macros name\n")
			return
		}

		macrosName = args[0]

		c := config.GetConfig()
		macros := c.GetMacros(macrosName)

		if macros == nil {
			color.Danger.Println("Macros not found")
			return
		}

		for _, v := range macros.List {
			cml := strings.Split(v, " ")
			if serviceName := c.FindService(cml[0]); serviceName != nil {
				serviceCmd.Run(cmd, cml)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(macrosCmd)
}
