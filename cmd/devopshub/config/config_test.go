package configcmd_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	configcmd "github.com/devopshub-org/devopshub-cli/cmd/devopshub/config"
)

func TestConfigCommands(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "devopshub-config-cmd-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	t.Setenv("HOME", tmpDir)
	t.Setenv("USERPROFILE", tmpDir)

	// Test init
	cmd := configcmd.NewConfigCmd()
	var outBuf bytes.Buffer
	cmd.SetOut(&outBuf)
	cmd.SetArgs([]string{"init"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("config init failed: %v", err)
	}
	if !strings.Contains(outBuf.String(), "Successfully initialized") {
		t.Errorf("unexpected init output: %s", outBuf.String())
	}

	// Test set
	cmdSet := configcmd.NewConfigCmd()
	var outSetBuf bytes.Buffer
	cmdSet.SetOut(&outSetBuf)
	cmdSet.SetArgs([]string{"set", "default_namespace", "my-app-ns"})
	if err := cmdSet.Execute(); err != nil {
		t.Fatalf("config set failed: %v", err)
	}

	// Test get
	cmdGet := configcmd.NewConfigCmd()
	var outGetBuf bytes.Buffer
	cmdGet.SetOut(&outGetBuf)
	cmdGet.SetArgs([]string{"get", "default_namespace"})
	if err := cmdGet.Execute(); err != nil {
		t.Fatalf("config get failed: %v", err)
	}
	if strings.TrimSpace(outGetBuf.String()) != "my-app-ns" {
		t.Errorf("expected 'my-app-ns', got '%s'", strings.TrimSpace(outGetBuf.String()))
	}
}
