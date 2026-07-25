package configcmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/devopshub-org/devopshub-cli/internal/config"
)

// NewConfigCmd creates the parent 'config' command and its subcommands.
func NewConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage devopshub CLI configuration",
		Long:  "View, initialize, set, and get values in the ~/.devopshub/config.yaml configuration file.",
	}

	cmd.AddCommand(
		newViewCmd(),
		newInitCmd(),
		newSetCmd(),
		newGetCmd(),
	)

	return cmd
}

func newViewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "view",
		Short: "Display current configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.InitConfig()
			if err != nil {
				return fmt.Errorf("failed to load configuration: %w", err)
			}

			cmd.Println("Configuration (~/.devopshub/config.yaml):")
			cmd.Printf("  Default Namespace : %s\n", cfg.DefaultNamespace)
			cmd.Printf("  Kubeconfig        : %s\n", cfg.Kubeconfig)
			cmd.Printf("  Output Format     : %s\n", cfg.OutputFormat)
			cmd.Println("  Cloud Profiles    :")
			for p, profile := range cfg.CloudProfiles {
				cmd.Printf("    %s: %s\n", p, profile)
			}

			return nil
		},
	}
}

func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Create default config file if it does not exist",
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile, err := config.GetConfigFile()
			if err != nil {
				return err
			}

			defaults := config.DefaultConfig()
			if err := config.WriteConfig(&defaults); err != nil {
				return fmt.Errorf("failed to initialize config: %w", err)
			}

			cmd.Printf("Successfully initialized configuration at %s\n", configFile)
			return nil
		},
	}
}

func newSetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a configuration value",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			key, value := args[0], args[1]

			cfg, err := config.InitConfig()
			if err != nil {
				cfg = &config.Config{CloudProfiles: make(map[string]string)}
			}

			switch key {
			case "default_namespace", "namespace":
				cfg.DefaultNamespace = value
			case "kubeconfig":
				cfg.Kubeconfig = value
			case "output_format", "format":
				cfg.OutputFormat = value
			default:
				return fmt.Errorf("unknown config key '%s' (valid: default_namespace, kubeconfig, output_format)", key)
			}

			if err := config.WriteConfig(cfg); err != nil {
				return fmt.Errorf("failed to save configuration: %w", err)
			}

			cmd.Printf("Set %s = %s\n", key, value)
			return nil
		},
	}
}

func newGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get <key>",
		Short: "Get a configuration value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]

			cfg, err := config.InitConfig()
			if err != nil {
				return fmt.Errorf("failed to load configuration: %w", err)
			}

			switch key {
			case "default_namespace", "namespace":
				cmd.Println(cfg.DefaultNamespace)
			case "kubeconfig":
				cmd.Println(cfg.Kubeconfig)
			case "output_format", "format":
				cmd.Println(cfg.OutputFormat)
			default:
				if val, exists := cfg.CloudProfiles[key]; exists {
					cmd.Println(val)
				} else {
					return fmt.Errorf("unknown config key '%s'", key)
				}
			}

			return nil
		},
	}
}
