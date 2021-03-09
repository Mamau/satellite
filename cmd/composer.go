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
		//composer config $(composer-repository) $(username) $(token);
		composer := entity.Composer{
			Version: composerVersion,
			WorkDir: "/home/www-data",
			HomeDir: "/home/www-data",
			Args:    append([]string{"composer"}, args...),
		}

		libs.RunCommandAtPTY(Docker(composer.CollectCommand()))
	},
}

func init() {
	rootCmd.AddCommand(composerCmd)
	composerCmd.Flags().StringVarP(&composerVersion, "version", "v", "2", "starter composer -v \"1.9\"")
}
