package commands

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"satellite/internal/entity"
	"strings"

	"satellite/pkg"

	"satellite/internal/config"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

type CommandRunner interface {
	Run() error
}

var rootCmd = &cobra.Command{
	Use:   "sat",
	Short: "All command",
	Long:  "Show all command",
}

func Docker(strategy entity.Runner, args []string) *exec.Cmd {
	replacedEnv := pkg.ReplaceEnvVariables(strategy.ToCommand(args))
	replacedPwd := pkg.ReplaceInternalVariables("\\$(\\(pwd\\))", pkg.GetPwd(), replacedEnv)
	replaceGateWay := getReplaceGateWay(replacedPwd)

	cmd := strategy.GetExecCommand()
	compose_2 := exec.Command("docker", "compose")
	cmdName, collectionAttributes, err := checkDockerService(cmd, replaceGateWay, exec.LookPath, compose_2)
	if err != nil {
		color.Red.Println(err)
		os.Exit(1)
	}

	dcCommand := exec.Command(cmdName, collectionAttributes...)
	color.Info.Printf("Running command: %v\n", dcCommand.String())
	return dcCommand
}

func checkDockerService(
	cmdName string,
	collectionAttributes []string,
	lookPath func(file string) (string, error),
	command CommandRunner,
) (string, []string, error) {
	if _, err := lookPath(cmdName); err == nil {
		return cmdName, collectionAttributes, nil
	}
	color.Warn.Printf("You have no %s.\n", cmdName)
	if cmdName == string(entity.DOCKER) {
		return "", nil, fmt.Errorf("%s not found", cmdName)
	}
	color.Warn.Println("Checking for docker compose 2nd version...")

	if err := command.Run(); err != nil {
		return "", nil, fmt.Errorf("oops... you need to install %s", cmdName)
	}

	prependedAttrs := append([]string{"compose"}, collectionAttributes...)

	return "docker", prependedAttrs, nil
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
