package cluster_test

import (
	"testing"

	"github.com/devopshub-org/devopshub-cli/cmd/devopshub/cluster"
)

func TestNewClusterCmd(t *testing.T) {
	cmd := cluster.NewClusterCmd()
	if cmd.Use != "cluster" {
		t.Errorf("expected command use 'cluster', got '%s'", cmd.Use)
	}
}
