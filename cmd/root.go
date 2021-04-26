package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/mamau/satellite/config"

	"github.com/gookit/color"
	"github.com/mamau/satellite/libs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "satlt",
	Short: "All command",
	Long:  "Show all command",
}

const commandName = "docker"

func Docker(dc Runnable) *exec.Cmd {
	replacedEnv := libs.ReplaceEnvVariables(dc.CollectCommand())
	replacedPwd := libs.ReplacePwdVariable(replacedEnv)
	dcCommand := exec.Command(commandName, replacedPwd...)
	color.Info.Printf("Running command: %v\n", dcCommand.String())
	return dcCommand
}

func Execute() {
	rc := getRunnableCommand()
	ac := getAvailableCommands()

	if _, has := libs.Find(ac, rc); has == false {
		c := config.GetConfig()
		if _, hasService := libs.Find(c.GetServices(), rc); hasService {
			serviceCmd.Run(rootCmd, os.Args[1:])
		} else {
			defaultExec()
		}
	} else {
		defaultExec()
	}
}

func defaultExec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getRunnableCommand() string {
	if isSet := libs.IndexExists(os.Args, 1); isSet {
		return os.Args[1]
	}
	return ""
}

func getAvailableCommands() []string {
	var availableCommands []string
	for _, v := range rootCmd.Commands() {
		availableCommands = append(availableCommands, v.Name())
	}
	return availableCommands
}
