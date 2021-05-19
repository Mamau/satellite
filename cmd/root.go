package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mamau/satellite/strategy"

	"github.com/mamau/satellite/config"

	"github.com/gookit/color"
	"github.com/mamau/satellite/libs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sat",
	Short: "All command",
	Long:  "Show all command",
}

const commandName = "docker"

func Docker(strategy strategy.Strategy) *exec.Cmd {
	replacedEnv := libs.ReplaceEnvVariables(strategy.ToCommand())
	replacedPwd := libs.ReplaceInternalVariables("\\$(\\(pwd\\))", libs.GetPwd(), replacedEnv)
	replaceGateWay := getReplaceGateWay(replacedPwd)

	dcCommand := exec.Command(commandName, replaceGateWay...)
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

func getReplaceGateWay(data []string) []string {
	from := "\\$(\\(gatewayHost\\))"
	r := regexp.MustCompile(from)
	if found := r.Find([]byte(strings.Join(data, " "))); found == nil {
		return data
	}

	inspectData := libs.DockerExec([]string{"network", "inspect", "bridge"})
	host := libs.RetrieveGatewayHost(inspectData)
	return libs.ReplaceInternalVariables(from, host, data)
}
