package main

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/devopshub-org/devopshub-cli/internal/output"
)

// rootCmd encapsulates the Cobra root command and shared state that must
// be accessible to all subcommands (flags, output printer, logger).
type rootCmd struct {
	cmd     *cobra.Command
	printer *output.Printer

	// Configurable writers. Defaults to os.Stdout/os.Stderr.
	// Tests override these to capture output for assertions.
	out    io.Writer
	errOut io.Writer

	// Persistent flag values — populated by Cobra before any RunE executes.
	noColor bool
	jsonOut bool
	debug   bool
	quiet   bool
}

// newRootCmd constructs the root command, wires persistent flags, and
// registers all subcommands. This is the sole entry point from main().
func newRootCmd() *rootCmd {
	r := &rootCmd{
		out:    os.Stdout,
		errOut: os.Stderr,
	}

	r.cmd = &cobra.Command{
		Use:   "devopshub",
		Short: "The open-source DevOps toolkit for engineers",
		Long: `DevOpsHub CLI — The world's best open-source DevOps toolkit.

A powerful, extensible command-line tool designed to streamline DevOps
workflows, manage cloud infrastructure, inspect Kubernetes clusters, and
automate routine operational tasks.

Documentation: https://github.com/devopshub-org/devopshub-cli
Issues:        https://github.com/devopshub-org/devopshub-cli/issues`,
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			// Re-initialize the printer with the flag values after Cobra has
			// parsed them. This runs before every subcommand's RunE.
			// Writers (r.out, r.errOut) are used so tests can inject buffers.
			r.printer = output.New(
				output.WithWriter(r.out),
				output.WithErrWriter(r.errOut),
				output.WithNoColor(r.noColor),
				output.WithJSON(r.jsonOut),
				output.WithQuiet(r.quiet),
			)

			// Configure zerolog level.
			level := zerolog.InfoLevel
			if r.debug {
				level = zerolog.DebugLevel
			}
			zerolog.SetGlobalLevel(level)
			_ = log.Logger // satisfy the import

			return nil
		},
	}

	// Persistent flags are inherited by all subcommands.
	pf := r.cmd.PersistentFlags()
	pf.BoolVar(&r.noColor, "no-color", false, "Disable ANSI color output")
	pf.BoolVarP(&r.jsonOut, "json", "j", false, "Output results as JSON")
	pf.BoolVar(&r.debug, "debug", false, "Enable debug logging")
	pf.BoolVarP(&r.quiet, "quiet", "q", false, "Suppress informational output")

	// Initialize printer with defaults (before flags are parsed).
	r.printer = output.New(
		output.WithWriter(r.out),
		output.WithErrWriter(r.errOut),
	)

	// Register all subcommands.
	registerSubcommands(r)

	return r
}

// Execute runs the root command. Called from main().
func (r *rootCmd) Execute() error {
	return r.cmd.Execute()
}
