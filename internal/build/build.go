// Package build provides version and build metadata injected at link time.
// All fields are populated via -ldflags during the release build process
// (GoReleaser compatible). Defaults to meaningful "dev" values for local builds.
package build

import "runtime"

// These variables are set at build time via -ldflags:
//
//	-X github.com/devopshub-org/devopshub-cli/internal/build.Version=v0.1.0
//	-X github.com/devopshub-org/devopshub-cli/internal/build.Commit=abc1234
//	-X github.com/devopshub-org/devopshub-cli/internal/build.Date=2024-01-01T00:00:00Z
//	-X github.com/devopshub-org/devopshub-cli/internal/build.BuiltBy=goreleaser
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
	BuiltBy = "source"
)

// Info holds all version and build metadata for the current binary.
type Info struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	Date      string `json:"date"`
	BuiltBy   string `json:"builtBy"`
	GoVersion string `json:"goVersion"`
	Platform  string `json:"platform"`
}

// Get returns the current build information.
func Get() Info {
	return Info{
		Version:   Version,
		Commit:    Commit,
		Date:      Date,
		BuiltBy:   BuiltBy,
		GoVersion: runtime.Version(),
		Platform:  runtime.GOOS + "/" + runtime.GOARCH,
	}
}
