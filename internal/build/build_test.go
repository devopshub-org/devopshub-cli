package build_test

import (
	"runtime"
	"strings"
	"testing"

	"github.com/devopshub-org/devopshub-cli/internal/build"
)

func TestGet_ReturnsInfo(t *testing.T) {
	info := build.Get()

	if info.Version == "" {
		t.Error("Version must not be empty")
	}
	if info.Commit == "" {
		t.Error("Commit must not be empty")
	}
	if info.Date == "" {
		t.Error("Date must not be empty")
	}
	if info.BuiltBy == "" {
		t.Error("BuiltBy must not be empty")
	}
	if info.GoVersion == "" {
		t.Error("GoVersion must not be empty")
	}
	if info.Platform == "" {
		t.Error("Platform must not be empty")
	}
}

func TestGet_DefaultsArePopulated(t *testing.T) {
	// Package-level defaults are sensible non-empty strings.
	info := build.Get()

	if info.Version != build.Version {
		t.Errorf("Get().Version = %q; want %q", info.Version, build.Version)
	}
	if info.Commit != build.Commit {
		t.Errorf("Get().Commit = %q; want %q", info.Commit, build.Commit)
	}
}

func TestGet_GoVersionMatchesRuntime(t *testing.T) {
	info := build.Get()
	if info.GoVersion != runtime.Version() {
		t.Errorf("GoVersion = %q; want %q", info.GoVersion, runtime.Version())
	}
}

func TestGet_PlatformFormat(t *testing.T) {
	info := build.Get()
	// Platform must be in the form "os/arch".
	if !strings.Contains(info.Platform, "/") {
		t.Errorf("Platform %q does not contain '/'; expected GOOS/GOARCH format", info.Platform)
	}
}
