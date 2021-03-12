package cmd

import (
	"github.com/gookit/color"
	"github.com/mamau/starter/entity"
	"github.com/mamau/starter/libs"
	"github.com/spf13/cobra"
)

var composerVersion string
var composerCmd = &cobra.Command{
	Use:   "composer",
	Short: "composer tool",
	Long:  "use it for interaction with composer",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("Start composer")

		if len(args) == 0 {
			args = []string{"--version"}
		}

		composer := entity.Composer{
			Command: entity.Command{
				Version:       composerVersion,
				Image:         "composer",
				HomeDir:       "/home/www-data",
				Args:          args,
				ConfigCommand: libs.GetConfig().GetComposer(),
			},
		}

		libs.RunCommandAtPTY(Docker(&composer))
	},
}

func init() {
	rootCmd.AddCommand(composerCmd)
	composerCmd.Flags().StringVarP(&composerVersion, "version", "v", "2", "starter composer -v \"1.9\"")
}
