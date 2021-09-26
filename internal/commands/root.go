package commands

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"satellite/internal/entity"

	"satellite/pkg"

	"satellite/internal/config"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sat",
	Short: "All command",
	Long:  "Show all command",
}

func Docker(strategy entity.Runner, args []string) *exec.Cmd {
	replacedEnv := pkg.ReplaceEnvVariables(strategy.ToCommand(args))
	replacedPwd := pkg.ReplaceInternalVariables("\\$(\\(pwd\\))", pkg.GetPwd(), replacedEnv)
	replaceGateWay := getReplaceGateWay(replacedPwd)

	dcCommand := exec.Command(strategy.GetExecCommand(), replaceGateWay...)
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
				serviceName := cmd.Name()
				color.Cyan.Printf("Start %s\n", serviceName)

				s := config.GetConfig().FindService(serviceName)

				pkg.RunCommandAtPTY(Docker(s, args))
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
