package plugin_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/devopshub-org/devopshub-cli/pkg/plugin"
)

func TestPluginDiscovery(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "devopshub-plugin-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	pluginDir := filepath.Join(tmpDir, ".devopshub", "plugins")
	if err := os.MkdirAll(pluginDir, 0755); err != nil {
		t.Fatalf("failed to create plugin dir: %v", err)
	}

	// Create dummy plugin file
	dummyPlugin := filepath.Join(pluginDir, "devopshub-testplugin")
	if err := os.WriteFile(dummyPlugin, []byte("#!/bin/sh\necho hello"), 0755); err != nil {
		t.Fatalf("failed to write dummy plugin: %v", err)
	}

	t.Setenv("HOME", tmpDir)
	t.Setenv("USERPROFILE", tmpDir)

	mgr, err := plugin.NewManager()
	if err != nil {
		t.Fatalf("failed to create manager: %v", err)
	}

	plugins, err := mgr.DiscoverPlugins()
	if err != nil {
		t.Fatalf("failed to discover plugins: %v", err)
	}

	found := false
	for _, p := range plugins {
		if p.Name == "testplugin" {
			found = true
			break
		}
	}

	if !found {
		t.Error("expected to discover 'testplugin'")
	}
}
