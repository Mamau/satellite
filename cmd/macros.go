package cmd

import (
	"strings"

	"github.com/gookit/color"

	"github.com/mamau/starter/config"

	"github.com/spf13/cobra"
)

var macrosCmd = &cobra.Command{
	Use:   "macros",
	Short: "Run group of commands",
	Long:  "Run group of commands",
	Run: func(cmd *cobra.Command, args []string) {
		var cl []string
		c := config.GetConfig()
		for _, v := range c.GetMacrosGroup() {
			cml := strings.Split(v, " ")
			for _, vv := range rootCmd.Commands() {
				if cml[0] == vv.Name() {
					cl = append(cl, cml[0])
					vv.Run(cmd, cml[1:])
				}
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
