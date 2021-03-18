package cmd

import (
	"github.com/gookit/color"
	"github.com/mamau/starter/entity"
	"github.com/mamau/starter/libs"

	"github.com/spf13/cobra"
)

var gulpCmd = &cobra.Command{
	Use:   "gulp",
	Short: "gulp tool",
	Long:  "use it for interaction with gulp",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("Start gulp")

		if len(args) == 0 {
			args = []string{"--version"}
		}

		gulp := entity.NewGulp(args)
		libs.RunCommandAtPTY(Docker(gulp))
	},
}

func init() {
	rootCmd.AddCommand(gulpCmd)
}
