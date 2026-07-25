package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Plugin struct {
	Name string
	Path string
}

type Manager struct {
	CustomDir string
}

func NewManager() (*Manager, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("unable to determine user home directory: %w", err)
	}
	return &Manager{
		CustomDir: filepath.Join(home, ".devopshub", "plugins"),
	}, nil
}

// DiscoverPlugins scans PATH and ~/.devopshub/plugins for devopshub-* binaries.
func (m *Manager) DiscoverPlugins() ([]Plugin, error) {
	plugins := make(map[string]string)

	// Scan custom plugin directory first
	if entries, err := os.ReadDir(m.CustomDir); err == nil {
		for _, e := range entries {
			if !e.IsDir() && strings.HasPrefix(e.Name(), "devopshub-") {
				name := strings.TrimPrefix(e.Name(), "devopshub-")
				name = strings.TrimSuffix(name, ".exe")
				plugins[name] = filepath.Join(m.CustomDir, e.Name())
			}
		}
	}

	// Scan PATH
	pathEnv := os.Getenv("PATH")
	for _, dir := range filepath.SplitList(pathEnv) {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if !e.IsDir() && strings.HasPrefix(e.Name(), "devopshub-") {
				name := strings.TrimPrefix(e.Name(), "devopshub-")
				name = strings.TrimSuffix(name, ".exe")
				if _, exists := plugins[name]; !exists {
					plugins[name] = filepath.Join(dir, e.Name())
				}
			}
		}
	}

	result := make([]Plugin, 0, len(plugins))
	for name, path := range plugins {
		result = append(result, Plugin{Name: name, Path: path})
	}

	return result, nil
}

// ExecutePlugin runs the plugin binary passing along args.
func (m *Manager) ExecutePlugin(pluginPath string, args []string) error {
	cmd := exec.Command(pluginPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	return cmd.Run()
}
