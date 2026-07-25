package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/devopshub-org/devopshub-cli/internal/config"
)

func TestDefaultConfig(t *testing.T) {
	cfg := config.DefaultConfig()
	if cfg.DefaultNamespace != "default" {
		t.Errorf("expected default namespace 'default', got '%s'", cfg.DefaultNamespace)
	}
	if cfg.OutputFormat != "table" {
		t.Errorf("expected default output format 'table', got '%s'", cfg.OutputFormat)
	}
}

func TestGetConfigDir(t *testing.T) {
	dir, err := config.GetConfigDir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if dir == "" {
		t.Error("expected non-empty config dir")
	}
}

func TestWriteAndInitConfig(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "devopshub-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Override HOME/USERPROFILE for testing
	t.Setenv("HOME", tmpDir)
	t.Setenv("USERPROFILE", tmpDir)

	cfg := &config.Config{
		DefaultNamespace: "test-ns",
		Kubeconfig:       "/path/to/kubeconfig",
		OutputFormat:     "json",
		CloudProfiles: map[string]string{
			"aws": "production",
		},
	}

	if err := config.WriteConfig(cfg); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	configFile := filepath.Join(tmpDir, ".devopshub", "config.yaml")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		t.Errorf("config file was not created at expected path: %s", configFile)
	}

	loaded, err := config.InitConfig()
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}

	if loaded.DefaultNamespace != "test-ns" {
		t.Errorf("expected 'test-ns', got '%s'", loaded.DefaultNamespace)
	}
	if loaded.OutputFormat != "json" {
		t.Errorf("expected 'json', got '%s'", loaded.OutputFormat)
	}
}
