package main

import (
	"github.com/spf13/cobra"
)

// completionCmd implements `devopshub completion`.
type completionCmd struct {
	cmd  *cobra.Command
	root *rootCmd
}

// newCompletionCmd constructs the shell completion subcommand.
// Cobra's built-in completion generation handles all shells.
func newCompletionCmd(root *rootCmd) *completionCmd {
	c := &completionCmd{root: root}

	c.cmd = &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate shell completion scripts",
		Long: `Generate shell completion scripts for DevOpsHub CLI.

To load completions:

  Bash:
    $ source <(devopshub completion bash)
    # To persist across sessions, add to ~/.bashrc:
    $ devopshub completion bash > /etc/bash_completion.d/devopshub

  Zsh:
    $ source <(devopshub completion zsh)
    # To persist across sessions, add to ~/.zshrc:
    $ devopshub completion zsh > "${fpath[1]}/_devopshub"

  Fish:
    $ devopshub completion fish | source
    # To persist across sessions:
    $ devopshub completion fish > ~/.config/fish/completions/devopshub.fish

  PowerShell:
    PS> devopshub completion powershell | Out-String | Invoke-Expression
    # To load for every new session, run and then add to your profile:
    PS> devopshub completion powershell > devopshub.ps1`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE:                  c.run,
	}

	return c
}

func (c *completionCmd) run(cmd *cobra.Command, args []string) error {
	switch args[0] {
	case "bash":
		return cmd.Root().GenBashCompletion(cmd.OutOrStdout())
	case "zsh":
		return cmd.Root().GenZshCompletion(cmd.OutOrStdout())
	case "fish":
		return cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
	case "powershell":
		return cmd.Root().GenPowerShellCompletionWithDesc(cmd.OutOrStdout())
	}
	// This branch is unreachable because cobra.OnlyValidArgs guards it.
	return nil
}
