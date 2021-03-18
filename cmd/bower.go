package cmd

import (
	"github.com/gookit/color"
	"github.com/mamau/starter/entity"
	"github.com/mamau/starter/libs"

	"github.com/spf13/cobra"
)

var bowerCmd = &cobra.Command{
	Use:   "bower",
	Short: "bower tool",
	Long:  "use it for interaction with bower",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("Start bower")

		if len(args) == 0 {
			args = []string{"--version"}
		}

		bower := entity.NewBower(args)
		libs.RunCommandAtPTY(Docker(bower))
	},
}

func init() {
	rootCmd.AddCommand(bowerCmd)
}
