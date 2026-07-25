package cloud

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ProviderStatus struct {
	Provider      string
	Authenticated bool
	Identity      string
	Details       string
}

type AuthHelper struct{}

func NewAuthHelper() *AuthHelper {
	return &AuthHelper{}
}

func (h *AuthHelper) CheckAWS(ctx context.Context) ProviderStatus {
	// Check AWS environment or credentials file
	home, err := os.UserHomeDir()
	credPath := ""
	if err == nil {
		credPath = filepath.Join(home, ".aws", "credentials")
	}

	hasEnv := os.Getenv("AWS_ACCESS_KEY_ID") != "" || os.Getenv("AWS_PROFILE") != ""
	_, hasFile := os.Stat(credPath)

	if !hasEnv && os.IsNotExist(hasFile) {
		return ProviderStatus{
			Provider:      "AWS",
			Authenticated: false,
			Identity:      "None",
			Details:       "No ~/.aws/credentials file or AWS environment variables found",
		}
	}

	// Try running aws sts get-caller-identity if available
	out, err := exec.CommandContext(ctx, "aws", "sts", "get-caller-identity", "--output", "json").Output()
	if err == nil {
		return ProviderStatus{
			Provider:      "AWS",
			Authenticated: true,
			Identity:      "Active (CLI)",
			Details:       strings.TrimSpace(string(out)),
		}
	}

	return ProviderStatus{
		Provider:      "AWS",
		Authenticated: true,
		Identity:      "Credentials Present",
		Details:       fmt.Sprintf("Credentials file present at %s", credPath),
	}
}

func (h *AuthHelper) CheckGCP(ctx context.Context) ProviderStatus {
	home, err := os.UserHomeDir()
	adcPath := ""
	if err == nil {
		adcPath = filepath.Join(home, ".config", "gcloud", "application_default_credentials.json")
	}

	hasEnv := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") != ""
	_, hasFile := os.Stat(adcPath)

	out, err := exec.CommandContext(ctx, "gcloud", "auth", "list", "--format=json").Output()
	if err == nil {
		return ProviderStatus{
			Provider:      "GCP",
			Authenticated: true,
			Identity:      "Active (gcloud CLI)",
			Details:       strings.TrimSpace(string(out)),
		}
	}

	if hasEnv || !os.IsNotExist(hasFile) {
		return ProviderStatus{
			Provider:      "GCP",
			Authenticated: true,
			Identity:      "Application Default Credentials Present",
			Details:       "GCP Application Default Credentials found",
		}
	}

	return ProviderStatus{
		Provider:      "GCP",
		Authenticated: false,
		Identity:      "None",
		Details:       "No gcloud credentials or GOOGLE_APPLICATION_CREDENTIALS found",
	}
}

func (h *AuthHelper) CheckAzure(ctx context.Context) ProviderStatus {
	home, err := os.UserHomeDir()
	azPath := ""
	if err == nil {
		azPath = filepath.Join(home, ".azure", "azureProfile.json")
	}

	out, err := exec.CommandContext(ctx, "az", "account", "show", "--output", "json").Output()
	if err == nil {
		return ProviderStatus{
			Provider:      "Azure",
			Authenticated: true,
			Identity:      "Active (az CLI)",
			Details:       strings.TrimSpace(string(out)),
		}
	}

	if _, err := os.Stat(azPath); err == nil {
		return ProviderStatus{
			Provider:      "Azure",
			Authenticated: true,
			Identity:      "Azure Profile Present",
			Details:       "Azure Profile found at ~/.azure/azureProfile.json",
		}
	}

	return ProviderStatus{
		Provider:      "Azure",
		Authenticated: false,
		Identity:      "None",
		Details:       "No active az CLI login or ~/.azure profile found",
	}
}
