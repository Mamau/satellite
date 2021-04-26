package cmd

import (
	"os"
	"strings"
	"testing"
)

func TestGetRunnableCommand(t *testing.T) {
	os.Args = []string{"satlt", "php", "-v"}
	if cmd := getRunnableCommand(); cmd != "php" {
		t.Errorf("runnable command must be %q", "php")
	}

	os.Args = []string{"satlt"}
	if cmd := getRunnableCommand(); cmd != "" {
		t.Error("runnable command must be empty")
	}
}

func TestGetAvailableCommands(t *testing.T) {
	ac := getAvailableCommands()
	var availableCommands []string
	for _, v := range rootCmd.Commands() {
		availableCommands = append(availableCommands, v.Name())
	}

	if strings.Join(ac, " ") != strings.Join(availableCommands, " ") {
		t.Errorf("available command expected %q\n got %q\n", strings.Join(ac, " "), strings.Join(availableCommands, " "))
	}
}
