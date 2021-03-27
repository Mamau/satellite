package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gookit/color"
	"github.com/mamau/starter/libs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "starter",
	Short: "All command",
	Long:  "Show all command",
}

func Docker(dc Runnable) *exec.Cmd {
	mainArgs := []string{"run", "-ti"}
	dcCommand := exec.Command("docker", append(mainArgs, libs.ReplaceEnvVariables(dc.CollectCommand())...)...)
	color.Info.Printf("Running command: %v\n", dcCommand.String())
	return dcCommand
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
