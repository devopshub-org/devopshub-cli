package cloud_test

import (
	"context"
	"testing"

	"github.com/devopshub-org/devopshub-cli/pkg/cloud"
)

func TestCheckAWS(t *testing.T) {
	helper := cloud.NewAuthHelper()
	status := helper.CheckAWS(context.Background())
	if status.Provider != "AWS" {
		t.Errorf("expected provider 'AWS', got '%s'", status.Provider)
	}
}

func TestCheckGCP(t *testing.T) {
	helper := cloud.NewAuthHelper()
	status := helper.CheckGCP(context.Background())
	if status.Provider != "GCP" {
		t.Errorf("expected provider 'GCP', got '%s'", status.Provider)
	}
}

func TestCheckAzure(t *testing.T) {
	helper := cloud.NewAuthHelper()
	status := helper.CheckAzure(context.Background())
	if status.Provider != "Azure" {
		t.Errorf("expected provider 'Azure', got '%s'", status.Provider)
	}
}
