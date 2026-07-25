package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/devopshub-org/devopshub-cli/internal/build"
)

// executeRoot is a test helper that creates a fresh root command with isolated
// output buffers and executes it with the provided arguments.
// Writers are injected via rootCmd.out/errOut fields, which PersistentPreRunE
// picks up when rebuilding the printer after flag parsing.
func executeRoot(args []string) (stdout, stderr string, err error) {
	var outBuf, errBuf bytes.Buffer
	root := newRootCmd()
	root.out = &outBuf
	root.errOut = &errBuf
	root.cmd.SetOut(&outBuf)
	root.cmd.SetErr(&errBuf)
	root.cmd.SetArgs(args)
	err = root.cmd.Execute()
	return outBuf.String(), errBuf.String(), err
}

// --- Root command ---

func TestRootCmd_NoArgs_ShowsHelp(t *testing.T) {
	out, _, err := executeRoot(nil)
	if err != nil {
		t.Fatalf("root with no args returned error: %v", err)
	}
	if !strings.Contains(out, "DevOpsHub") && !strings.Contains(out, "Usage") {
		t.Errorf("root help should mention DevOpsHub or Usage\ngot: %s", out)
	}
}

func TestRootCmd_Help_NoError(t *testing.T) {
	_, _, err := executeRoot([]string{"--help"})
	if err != nil {
		t.Fatalf("--help returned error: %v", err)
	}
}

func TestRootCmd_UnknownCommand_ReturnsError(t *testing.T) {
	_, _, err := executeRoot([]string{"nonexistent"})
	if err == nil {
		t.Error("expected error for unknown command, got nil")
	}
}

func TestRootCmd_HasVersionSubcommand(t *testing.T) {
	root := newRootCmd()
	var found bool
	for _, sub := range root.cmd.Commands() {
		if sub.Name() == "version" {
			found = true
			break
		}
	}
	if !found {
		t.Error("root command must have 'version' subcommand")
	}
}

func TestRootCmd_HasCheatSubcommand(t *testing.T) {
	root := newRootCmd()
	var found bool
	for _, sub := range root.cmd.Commands() {
		if sub.Name() == "cheat" {
			found = true
			break
		}
	}
	if !found {
		t.Error("root command must have 'cheat' subcommand")
	}
}

func TestRootCmd_HasCompletionSubcommand(t *testing.T) {
	root := newRootCmd()
	var found bool
	for _, sub := range root.cmd.Commands() {
		if sub.Name() == "completion" {
			found = true
			break
		}
	}
	if !found {
		t.Error("root command must have 'completion' subcommand")
	}
}

// --- Version command ---

func TestVersionCmd_DefaultOutput(t *testing.T) {
	out, _, err := executeRoot([]string{"version"})
	if err != nil {
		t.Fatalf("version returned error: %v", err)
	}
	if !strings.Contains(out, "devopshub") {
		t.Errorf("version output should contain 'devopshub'\ngot: %s", out)
	}
}

func TestVersionCmd_ContainsBuildVersion(t *testing.T) {
	out, _, err := executeRoot([]string{"version"})
	if err != nil {
		t.Fatalf("version returned error: %v", err)
	}
	// The default build.Version is "dev" in test builds.
	if !strings.Contains(out, build.Version) {
		t.Errorf("version output should contain version %q\ngot: %s", build.Version, out)
	}
}

func TestVersionCmd_JSONOutput_IsValidJSON(t *testing.T) {
	out, _, err := executeRoot([]string{"--json", "version"})
	if err != nil {
		t.Fatalf("version --json returned error: %v", err)
	}
	var info build.Info
	if jsonErr := json.Unmarshal([]byte(out), &info); jsonErr != nil {
		t.Errorf("version --json output is not valid JSON: %v\noutput: %s", jsonErr, out)
	}
}

func TestVersionCmd_JSONOutput_ContainsRequiredFields(t *testing.T) {
	out, _, err := executeRoot([]string{"--json", "version"})
	if err != nil {
		t.Fatalf("version --json returned error: %v", err)
	}
	for _, field := range []string{"version", "commit", "date", "goVersion", "platform"} {
		if !strings.Contains(out, field) {
			t.Errorf("version --json missing field %q\noutput: %s", field, out)
		}
	}
}

func TestVersionCmd_VerboseOutput(t *testing.T) {
	out, _, err := executeRoot([]string{"version", "--verbose"})
	if err != nil {
		t.Fatalf("version --verbose returned error: %v", err)
	}
	for _, want := range []string{"Version", "Commit", "Platform"} {
		if !strings.Contains(out, want) {
			t.Errorf("version --verbose missing %q\noutput: %s", want, out)
		}
	}
}

func TestVersionCmd_Alias(t *testing.T) {
	root := newRootCmd()
	for _, sub := range root.cmd.Commands() {
		if sub.Name() == "version" {
			for _, alias := range sub.Aliases {
				if alias == "v" {
					return
				}
			}
			t.Error("version command must have 'v' alias")
		}
	}
}

// --- Completion ---

func TestCompletionCmd_Bash_NoError(t *testing.T) {
	_, _, err := executeRoot([]string{"completion", "bash"})
	if err != nil {
		t.Fatalf("completion bash returned error: %v", err)
	}
}

func TestCompletionCmd_Zsh_NoError(t *testing.T) {
	_, _, err := executeRoot([]string{"completion", "zsh"})
	if err != nil {
		t.Fatalf("completion zsh returned error: %v", err)
	}
}

func TestCompletionCmd_PowerShell_NoError(t *testing.T) {
	_, _, err := executeRoot([]string{"completion", "powershell"})
	if err != nil {
		t.Fatalf("completion powershell returned error: %v", err)
	}
}

func TestCompletionCmd_InvalidShell_ReturnsError(t *testing.T) {
	_, _, err := executeRoot([]string{"completion", "nushell"})
	if err == nil {
		t.Error("expected error for unsupported shell 'nushell'")
	}
}

func TestCompletionCmd_NoArgs_ReturnsError(t *testing.T) {
	_, _, err := executeRoot([]string{"completion"})
	if err == nil {
		t.Error("expected error when no shell argument is provided")
	}
}
