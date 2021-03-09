package cmd

import (
	"github.com/gookit/color"
	"github.com/mamau/starter/entity"
	"github.com/mamau/starter/libs"

	"github.com/spf13/cobra"
)

var nodeForYarnVersion string
var yarnCmd = &cobra.Command{
	Use:   "yarn",
	Short: "yarn tool",
	Long:  "use it for interaction with yarn",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("Start yarn")

		if len(args) == 0 {
			args = []string{"--version"}
		}

		yarn := entity.Yarn{
			Version: nodeForYarnVersion,
			WorkDir: "/home/node",
			HomeDir: "/home/node",
			Args:    append([]string{"yarn"}, args...),
		}

		libs.RunCommandAtPTY(Docker(yarn.CollectCommand()))
	},
}

func init() {
	rootCmd.AddCommand(yarnCmd)
	yarnCmd.Flags().StringVarP(&nodeForYarnVersion, "version", "v", "10", "starter yarn -v \"14\"")
}
