package commands

import (
	"satellite/internal/updater"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Self update",
	Long:  "Self update to latest version",
	Run: func(cmd *cobra.Command, args []string) {
		updater := updater.NewSelfUpdater()
		updater.Update()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
