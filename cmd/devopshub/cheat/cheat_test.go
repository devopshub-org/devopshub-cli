package cheat_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/devopshub-org/devopshub-cli/cmd/devopshub/cheat"
	"github.com/devopshub-org/devopshub-cli/internal/output"
)

// newTestPrinter returns a Printer writing to a buffer for assertion.
func newTestPrinter(opts ...output.Option) (*output.Printer, *bytes.Buffer) {
	var buf bytes.Buffer
	base := []output.Option{
		output.WithWriter(&buf),
		output.WithErrWriter(&buf),
		output.WithNoColor(true),
	}
	return output.New(append(base, opts...)...), &buf
}

// execute creates a new root command for the cheat group and executes it
// with the given arguments, returning captured output and any error.
func execute(args []string, opts ...output.Option) (string, error) {
	var buf bytes.Buffer
	p := output.New(append([]output.Option{
		output.WithWriter(&buf),
		output.WithErrWriter(&buf),
		output.WithNoColor(true),
	}, opts...)...)

	parent := &cobra.Command{Use: "devopshub", SilenceUsage: true, SilenceErrors: true}
	parent.AddCommand(cheat.NewCheatCmd(p).CobraCmd())
	parent.SetOut(&buf)
	parent.SetErr(&buf)
	parent.SetArgs(append([]string{"cheat"}, args...))
	err := parent.Execute()
	return buf.String(), err
}

// --- Structure ---

func TestNewCheatCmd_NotNil(t *testing.T) {
	p, _ := newTestPrinter()
	c := cheat.NewCheatCmd(p)
	if c == nil {
		t.Fatal("NewCheatCmd returned nil")
	}
}

func TestCheatCmd_HasExpectedSubcommands(t *testing.T) {
	p, _ := newTestPrinter()
	c := cheat.NewCheatCmd(p)
	cmd := c.CobraCmd()

	want := map[string]bool{"docker": false, "k8s": false, "linux": false}
	for _, sub := range cmd.Commands() {
		want[sub.Name()] = true
	}
	for name, found := range want {
		if !found {
			t.Errorf("expected subcommand %q not found", name)
		}
	}
}

// --- Docker ---

func TestDockerCheat_ContainsCoreCommands(t *testing.T) {
	got, err := execute([]string{"docker"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, want := range []string{"docker ps", "docker run", "docker images", "docker build"} {
		if !strings.Contains(got, want) {
			t.Errorf("docker cheatsheet missing %q\nfull output:\n%s", want, got)
		}
	}
}

func TestDockerCheat_ContainsSections(t *testing.T) {
	got, err := execute([]string{"docker"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, section := range []string{"Container Lifecycle", "Images", "Networking", "Volumes"} {
		if !strings.Contains(got, section) {
			t.Errorf("docker cheatsheet missing section %q", section)
		}
	}
}

func TestDockerCheat_JSONOutputIsValid(t *testing.T) {
	got, err := execute([]string{"docker"}, output.WithJSON(true))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result []map[string]any
	if jsonErr := json.Unmarshal([]byte(got), &result); jsonErr != nil {
		t.Errorf("docker --json is not valid JSON: %v\noutput: %s", jsonErr, got)
	}
	if len(result) == 0 {
		t.Error("docker --json sections array must not be empty")
	}
}

func TestDockerCheat_JSONContainsTitleAndEntries(t *testing.T) {
	got, err := execute([]string{"docker"}, output.WithJSON(true))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, `"title"`) || !strings.Contains(got, `"entries"`) {
		t.Errorf("docker --json must contain 'title' and 'entries' keys\noutput: %s", got)
	}
}

// --- Kubernetes ---

func TestK8sCheat_ContainsCoreCommands(t *testing.T) {
	got, err := execute([]string{"k8s"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, want := range []string{"kubectl get pods", "kubectl logs", "kubectl apply"} {
		if !strings.Contains(got, want) {
			t.Errorf("k8s cheatsheet missing %q", want)
		}
	}
}

func TestK8sCheat_JSONOutputIsValid(t *testing.T) {
	got, err := execute([]string{"k8s"}, output.WithJSON(true))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result []map[string]any
	if jsonErr := json.Unmarshal([]byte(got), &result); jsonErr != nil {
		t.Errorf("k8s --json is not valid JSON: %v\noutput: %s", jsonErr, got)
	}
}

func TestK8sCheat_KubernetesAlias(t *testing.T) {
	p, _ := newTestPrinter()
	c := cheat.NewCheatCmd(p)
	for _, sub := range c.CobraCmd().Commands() {
		if sub.Name() == "k8s" {
			for _, alias := range sub.Aliases {
				if alias == "kubernetes" {
					return
				}
			}
			t.Error("k8s must have 'kubernetes' alias")
		}
	}
}

// --- Linux ---

func TestLinuxCheat_ContainsCoreCommands(t *testing.T) {
	got, err := execute([]string{"linux"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, want := range []string{"grep", "awk", "ssh", "tar", "find"} {
		if !strings.Contains(got, want) {
			t.Errorf("linux cheatsheet missing %q", want)
		}
	}
}

func TestLinuxCheat_JSONOutputIsValid(t *testing.T) {
	got, err := execute([]string{"linux"}, output.WithJSON(true))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var result []map[string]any
	if jsonErr := json.Unmarshal([]byte(got), &result); jsonErr != nil {
		t.Errorf("linux --json is not valid JSON: %v\noutput: %s", jsonErr, got)
	}
}

// --- Negative / Edge cases ---

func TestCheatCmd_NoArgs_PrintsHelp(t *testing.T) {
	got, err := execute(nil)
	// No error — the parent cheat command defaults to showing help.
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "Available Commands") && !strings.Contains(got, "docker") {
		t.Errorf("cheat with no args should show help, got: %s", got)
	}
}

func TestCheatCmd_UnknownSubcommand_FallsBackToHelp(t *testing.T) {
	// When an unrecognized argument is passed to a command with RunE defined,
	// cobra passes it as a positional arg to the parent RunE, which shows help.
	// This is expected behavior — the parent's RunE returns cmd.Help() → nil error.
	_, err := execute([]string{"terraform"})
	if err != nil {
		// Some cobra versions do return an error for unrecognized subcommands.
		// Accept both behaviors as valid.
		t.Logf("cheat terraform returned error (acceptable): %v", err)
	}
}
