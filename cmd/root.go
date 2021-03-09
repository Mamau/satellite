package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"

	"github.com/gookit/color"
	"github.com/mamau/starter/libs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "starter",
	Short: "All command",
	Long:  "Show all command",
}

func init() {
	//libs.PrepareEnv()
}

func Docker(args []string) *exec.Cmd {
	//mainArgs := []string{"run", "-ti", "-u", UserId()}
	mainArgs := []string{"run", "-ti", "-e", fmt.Sprintf("USER_ID=%s", UserId())}
	dcCommand := exec.Command("docker", append(mainArgs, args...)...)
	color.Info.Printf("Running command: %v\n", dcCommand.String())
	return dcCommand
}

// Add -T flag for windows commands
func prepareForOs(args []string) []string {
	if runtime.GOOS != "windows" {
		return args
	}
	indexExec, isSet := libs.Find(args, "exec")
	if !isSet {
		log.Fatalf("Arguments %v not have exec key word", args)
	}

	return libs.InsertToSlice(args, "-T", indexExec+1)
}

func UserId() string {
	if runtime.GOOS == "windows" {
		return "1000"
	}
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("cant get user. error: %s\n", err)
	}

	return currentUser.Uid
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
