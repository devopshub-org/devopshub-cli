// Package cheat implements the `devopshub cheat` command family.
// Each subcommand renders a curated, categorized cheatsheet for a specific
// technology to stdout. Content is static and embedded in the binary —
// no network access, no external files.
package cheat

import (
	"github.com/spf13/cobra"

	"github.com/devopshub-org/devopshub-cli/internal/output"
)

// Cmd wraps the Cobra command and shared printer for the cheat group.
type Cmd struct {
	cmd     *cobra.Command
	printer *output.Printer
}

// NewCheatCmd constructs the `cheat` parent command and registers all
// cheatsheet subcommands.
func NewCheatCmd(printer *output.Printer) *Cmd {
	c := &Cmd{printer: printer}

	c.cmd = &cobra.Command{
		Use:   "cheat",
		Short: "Interactive cheatsheets for DevOps tools",
		Long: `Display curated, categorized cheatsheets for popular DevOps tools.

All cheatsheets are embedded in the binary — no internet connection required.

Available cheatsheets:
  docker      Docker commands and best practices
  k8s         Kubernetes (kubectl) commands
  linux       Essential Linux commands for DevOps

Examples:
  devopshub cheat docker
  devopshub cheat k8s
  devopshub cheat linux
  devopshub cheat docker --json`,
		// Printing the help is the natural action when no subcommand is given.
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	c.cmd.AddCommand(
		newDockerCmd(printer).cmd,
		newK8sCmd(printer).cmd,
		newLinuxCmd(printer).cmd,
	)

	return c
}

// CobraCmd returns the underlying *cobra.Command.
// This is the public accessor for test and integration use.
func (c *Cmd) CobraCmd() *cobra.Command {
	return c.cmd
}
