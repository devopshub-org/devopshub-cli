package cloudcmd_test

import (
	"testing"

	cloudcmd "github.com/devopshub-org/devopshub-cli/cmd/devopshub/cloud"
)

func TestNewCloudCmd(t *testing.T) {
	cmd := cloudcmd.NewCloudCmd()
	if cmd.Use != "cloud" {
		t.Errorf("expected command use 'cloud', got '%s'", cmd.Use)
	}
}
