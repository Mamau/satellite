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

		gulp := entity.Gulp{
			WorkDir: "/home/node",
			HomeDir: "/home/node",
			Args:    args,
		}

		libs.RunCommandAtPTY(Docker(gulp.CollectCommand()))
	},
}

func init() {
	rootCmd.AddCommand(gulpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gulpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gulpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
