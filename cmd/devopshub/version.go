package main

import (
	"github.com/spf13/cobra"

	"github.com/devopshub-org/devopshub-cli/internal/build"
)

// versionCmd implements the `devopshub version` command.
type versionCmd struct {
	cmd     *cobra.Command
	root    *rootCmd
	verbose bool
}

// newVersionCmd constructs the version subcommand.
func newVersionCmd(root *rootCmd) *versionCmd {
	v := &versionCmd{root: root}

	v.cmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		Long: `Print the current version, commit hash, build date, and runtime
information for this DevOpsHub CLI binary.

Examples:
  devopshub version
  devopshub version --verbose
  devopshub version --json`,
		Args:    cobra.NoArgs,
		RunE:    v.run,
		Aliases: []string{"v"},
	}

	v.cmd.Flags().BoolVarP(&v.verbose, "verbose", "v", false, "Print full build metadata")

	return v
}

func (v *versionCmd) run(_ *cobra.Command, _ []string) error {
	info := build.Get()
	p := v.root.printer

	// JSON mode: emit structured JSON and return immediately.
	if v.root.jsonOut {
		return p.JSON(info)
	}

	// Verbose mode: print all build metadata.
	if v.verbose {
		p.Header("DevOpsHub CLI — Build Information")
		p.Println("")
		p.Printf("  Version   : %s\n", info.Version)
		p.Printf("  Commit    : %s\n", info.Commit)
		p.Printf("  Built     : %s\n", info.Date)
		p.Printf("  Built By  : %s\n", info.BuiltBy)
		p.Printf("  Go        : %s\n", info.GoVersion)
		p.Printf("  Platform  : %s\n", info.Platform)
		p.Println("")
		return nil
	}

	// Default: single-line version string (machine-parseable).
	p.Printf("devopshub version %s (%s) %s\n", info.Version, info.Commit, info.Platform)
	return nil
}
