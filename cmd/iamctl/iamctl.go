package main

// iamctl is the command line tool for iam platform.

import (
	"go_iam/internal/iamctl/cmd"
	"os"
)

func main() {
	command := cmd.NewDefaultIAMCtlCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
