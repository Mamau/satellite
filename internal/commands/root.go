package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"satellite/internal/entity"
	"strings"
	"syscall"

	"satellite/pkg"

	"satellite/internal/config"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

const FAILURE = 1

type CommandRunner interface {
	Run() error
}

func newCommandRunner(name string, arg ...string) CommandRunner {
	return exec.Command(name, arg...)
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
	cmdName, collectionAttributes, err := checkDockerService(cmd, replaceGateWay, exec.LookPath, newCommandRunner)
	if err != nil {
		color.Red.Println(err)
		os.Exit(FAILURE)
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		color.Info.Println("Stopping...")
		cancel()
	}()
	dcCommand := exec.CommandContext(ctx, cmdName, collectionAttributes...)

	color.Info.Printf("Running command: %v\n", dcCommand.String())
	return dcCommand
}

func checkDockerService(
	cmdName string,
	collectionAttributes []string,
	lookPath func(file string) (string, error),
	command func(name string, arg ...string) CommandRunner,
) (string, []string, error) {
	if _, err := lookPath(cmdName); err == nil {
		return cmdName, collectionAttributes, nil
	}
	color.Warn.Printf("You have no %s.\n", cmdName)
	if cmdName == string(entity.DOCKER) {
		return "", nil, fmt.Errorf("%s not found", cmdName)
	}
	color.Warn.Println("Checking for docker compose 2nd version...")

	cmd := command("docker", "compose")
	if err := cmd.Run(); err != nil {
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

				eCmd := Docker(s, args)
				eCmd.Stderr = os.Stderr
				eCmd.Stdout = os.Stdout
				eCmd.Stdin = os.Stdin

				if err := eCmd.Start(); err != nil {
					log.Fatalf("cmd.Start: %v", err)
				}
				if err := eCmd.Wait(); err != nil {
					if exiterr, ok := err.(*exec.ExitError); ok {

						if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
							if status.ExitStatus() == FAILURE {
								os.Exit(FAILURE)
							}
						}
					} else {
						log.Fatalf("cmd.Wait: %v", err)
					}
				}
			},
		})
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(FAILURE)
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
