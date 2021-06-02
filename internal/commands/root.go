package commands

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mamau/satellite/pkg"

	config2 "github.com/mamau/satellite/internal/config"

	strategy2 "github.com/mamau/satellite/internal/strategy"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sat",
	Short: "All command",
	Long:  "Show all command",
}

const commandName = "docker"

func Docker(strategy strategy2.Strategy) *exec.Cmd {
	replacedEnv := pkg.ReplaceEnvVariables(strategy.ToCommand())
	replacedPwd := pkg.ReplaceInternalVariables("\\$(\\(pwd\\))", pkg.GetPwd(), replacedEnv)
	replaceGateWay := getReplaceGateWay(replacedPwd)

	dcCommand := exec.Command(commandName, replaceGateWay...)
	color.Info.Printf("Running command: %v\n", dcCommand.String())
	return dcCommand
}

func Execute() {
	rc := getRunnableCommand()
	ac := getAvailableCommands()

	if _, has := pkg.Find(ac, rc); has == false {
		c := config2.GetConfig()
		if _, hasService := pkg.Find(c.GetServices(), rc); hasService {
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
	if isSet := pkg.IndexExists(os.Args, 1); isSet {
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

	inspectData := pkg.DockerExec([]string{"network", "inspect", "bridge"})
	host := pkg.RetrieveGatewayHost(inspectData)
	return pkg.ReplaceInternalVariables(from, host, data)
}
