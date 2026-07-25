package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	DefaultNamespace string            `mapstructure:"default_namespace" yaml:"default_namespace"`
	Kubeconfig       string            `mapstructure:"kubeconfig" yaml:"kubeconfig"`
	OutputFormat     string            `mapstructure:"output_format" yaml:"output_format"`
	CloudProfiles    map[string]string `mapstructure:"cloud_profiles" yaml:"cloud_profiles"`
}

// DefaultConfig returns the default configuration settings.
func DefaultConfig() Config {
	return Config{
		DefaultNamespace: "default",
		Kubeconfig:       "",
		OutputFormat:     "table",
		CloudProfiles: map[string]string{
			"aws":   "default",
			"gcp":   "default",
			"azure": "default",
		},
	}
}

// GetConfigDir returns the path to ~/.devopshub
func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to determine user home directory: %w", err)
	}
	return filepath.Join(home, ".devopshub"), nil
}

// GetConfigFile returns the path to ~/.devopshub/config.yaml
func GetConfigFile() (string, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.yaml"), nil
}

// InitConfig initializes viper and loads configuration.
func InitConfig() (*Config, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(dir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("DEVOPSHUB")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Set defaults
	defaults := DefaultConfig()
	viper.SetDefault("default_namespace", defaults.DefaultNamespace)
	viper.SetDefault("kubeconfig", defaults.Kubeconfig)
	viper.SetDefault("output_format", defaults.OutputFormat)
	viper.SetDefault("cloud_profiles", defaults.CloudProfiles)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &cfg, nil
}

// WriteConfig creates or updates ~/.devopshub/config.yaml
func WriteConfig(cfg *Config) error {
	dir, err := GetConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configFile, err := GetConfigFile()
	if err != nil {
		return err
	}

	viper.Set("default_namespace", cfg.DefaultNamespace)
	viper.Set("kubeconfig", cfg.Kubeconfig)
	viper.Set("output_format", cfg.OutputFormat)
	viper.Set("cloud_profiles", cfg.CloudProfiles)

	if err := viper.WriteConfigAs(configFile); err != nil {
		if os.IsExist(err) {
			return viper.WriteConfig()
		}
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
