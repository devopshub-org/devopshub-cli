package configcmd

import (
	"fmt"

	"github.com/devopshub-org/devopshub-cli/internal/config"
	"github.com/spf13/cobra"
)

// NewConfigCmd creates the 'config' command and its subcommands.
func NewConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage DevOpsHub CLI user configuration",
		Long:  "Initialize, view, and modify configuration stored at ~/.devopshub/config.yaml.",
	}

	cmd.AddCommand(
		newInitCmd(),
		newViewCmd(),
		newGetCmd(),
		newSetCmd(),
	)

	return cmd
}

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize default configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			defaults := config.DefaultConfig()
			if err := config.WriteConfig(&defaults); err != nil {
				return err
			}
			path, _ := config.GetConfigFile()
			fmt.Fprintf(cmd.OutOrStdout(), "Successfully initialized configuration at %s\n", path)
			return nil
		},
	}
}

func newViewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "view",
		Short: "Display active configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.InitConfig()
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Default Namespace: %s\n", cfg.DefaultNamespace)
			fmt.Fprintf(cmd.OutOrStdout(), "Kubeconfig:        %s\n", cfg.Kubeconfig)
			fmt.Fprintf(cmd.OutOrStdout(), "Output Format:     %s\n", cfg.OutputFormat)
			fmt.Fprintln(cmd.OutOrStdout(), "Cloud Profiles:")
			for provider, profile := range cfg.CloudProfiles {
				fmt.Fprintf(cmd.OutOrStdout(), "  %s: %s\n", provider, profile)
			}
			return nil
		},
	}
}

func newGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get [key]",
		Short: "Get a specific configuration setting",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.InitConfig()
			if err != nil {
				return err
			}
			key := args[0]
			switch key {
			case "default_namespace", "namespace":
				fmt.Fprintln(cmd.OutOrStdout(), cfg.DefaultNamespace)
			case "kubeconfig":
				fmt.Fprintln(cmd.OutOrStdout(), cfg.Kubeconfig)
			case "output_format", "format":
				fmt.Fprintln(cmd.OutOrStdout(), cfg.OutputFormat)
			default:
				return fmt.Errorf("unknown configuration key: %s", key)
			}
			return nil
		},
	}
}

func newSetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set [key] [value]",
		Short: "Set a specific configuration setting",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.InitConfig()
			if err != nil {
				return err
			}
			key, val := args[0], args[1]
			switch key {
			case "default_namespace", "namespace":
				cfg.DefaultNamespace = val
			case "kubeconfig":
				cfg.Kubeconfig = val
			case "output_format", "format":
				cfg.OutputFormat = val
			default:
				return fmt.Errorf("unknown configuration key: %s", key)
			}
			if err := config.WriteConfig(cfg); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Updated %s to %s\n", key, val)
			return nil
		},
	}
}
