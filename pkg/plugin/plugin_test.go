package plugin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPluginDiscovery(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "devopshub-plugin-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	pluginDir := filepath.Join(tmpDir, ".devopshub", "plugins")
	if err := os.MkdirAll(pluginDir, 0o755); err != nil {
		t.Fatalf("failed to create plugin dir: %v", err)
	}

	// Create a mock executable plugin file
	dummyPlugin := filepath.Join(pluginDir, "devopshub-hello")
	if err := os.WriteFile(dummyPlugin, []byte("#!/bin/sh\necho hello"), 0o755); err != nil {
		t.Fatalf("failed to create dummy plugin: %v", err)
	}

	mgr := &Manager{CustomDir: pluginDir}
	plugins, err := mgr.DiscoverPlugins()
	if err != nil {
		t.Fatalf("failed to discover plugins: %v", err)
	}

	if len(plugins) != 1 {
		t.Errorf("expected 1 plugin, got %d", len(plugins))
	} else if plugins[0].Name != "hello" {
		t.Errorf("expected plugin name 'hello', got '%s'", plugins[0].Name)
	}
}
