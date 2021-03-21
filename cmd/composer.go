package cmd

import (
	"github.com/mamau/starter/libs"

	"github.com/gookit/color"
	"github.com/mamau/starter/entity"
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

		composer := entity.NewComposer(composerVersion, args)
		libs.RunCommandAtPTY(Docker(composer))
	},
}

func init() {
	rootCmd.AddCommand(composerCmd)
	composerCmd.Flags().StringVarP(&composerVersion, "version", "v", "2", "starter composer -v \"1.9\"")
}
