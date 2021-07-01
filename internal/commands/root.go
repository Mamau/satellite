package commands

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/mamau/satellite/pkg"

	"github.com/mamau/satellite/internal/config"

	"github.com/mamau/satellite/internal/strategy"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sat",
	Short: "All command",
	Long:  "Show all command",
}

const commandName = "docker"

func Docker(strategy strategy.Strategy) *exec.Cmd {
	replacedEnv := pkg.ReplaceEnvVariables(strategy.ToCommand())
	replacedPwd := pkg.ReplaceInternalVariables("\\$(\\(pwd\\))", pkg.GetPwd(), replacedEnv)
	replaceGateWay := getReplaceGateWay(replacedPwd)

	dcCommand := exec.Command(commandName, replaceGateWay...)
	color.Info.Printf("Running command: %v\n", dcCommand.String())
	return dcCommand
}

func InitServiceCommand() {
	c := config.GetConfig()
	for _, service := range c.GetServices() {
		rootCmd.AddCommand(&cobra.Command{
			Use:                service.Name,
			Short:              service.Description,
			Long:               service.Description,
			DisableFlagParsing: true,
			Run: func(cmd *cobra.Command, args []string) {
				//if len(args) < 1 {
				//	color.Red.Printf("You should pass service name")
				//	return
				//}

				serviceName := cmd.Name()
				color.Cyan.Printf("Start %s\n", serviceName)
				s := config.GetConfig().GetService(serviceName)
				strgy := determineStrategy(s, args)
				//fmt.Println(len(args), "------")

				pkg.RunCommandAtPTY(Docker(strgy))
			},
		})
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
