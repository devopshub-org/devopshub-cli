// Package main is the entry point for the DevOpsHub CLI.
// It delegates all work to the root Cobra command defined in root.go.
package main

import (
	"os"

	"github.com/devopshub-org/devopshub-cli/cmd/devopshub/cheat"
)

func main() {
	if err := newRootCmd().Execute(); err != nil {
		// Cobra already printed the error. Exit with a non-zero code.
		os.Exit(1)
	}
}

// registerSubcommands wires all top-level subcommands into the root command.
// Adding a new top-level command only requires a single line here.
func registerSubcommands(root *rootCmd) {
	root.cmd.AddCommand(
		newVersionCmd(root).cmd,
		newCompletionCmd(root).cmd,
		cheat.NewCheatCmd(root.printer).CobraCmd(),
	)
}
