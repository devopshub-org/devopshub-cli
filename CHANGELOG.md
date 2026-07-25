# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [Unreleased]

## [0.1.0] - 2026-07-25

### Added

**CLI Foundation**
- Root Cobra command with persistent flags: `--no-color`, `--json`, `--debug`, `--quiet`
- `version` command: single-line, `--verbose`, and `--json` output modes
- `completion` command: bash, zsh, fish, and PowerShell shell completion generation
- `cheat docker` command: 50+ Docker commands in 7 categorized sections
- `cheat k8s` command: 60+ kubectl commands in 9 categorized sections (aliases: `kubernetes`, `kubectl`)
- `cheat linux` command: 100+ Linux commands in 8 categorized sections (aliases: `unix`, `shell`)
- All cheatsheets support `--json` for machine-readable output

**Internal Packages**
- `internal/build` — Version metadata struct with ldflags injection support
- `internal/output` — Terminal output primitives (`Success`, `Error`, `Info`, `Warning`, `Header`, `Table`, `JSON`, `Command`, `Section`). Single source of truth for all user-facing rendering. Supports `--no-color`, `--json`, `--quiet`

**Build System**
- Production Makefile: `build`, `build-all`, `test`, `test-coverage`, `lint`, `fmt`, `vet`, `tidy`, `install`, `release-snapshot`, `clean`, `help`
- Cross-platform build matrix: Linux, macOS, Windows (amd64/arm64)
- Version injection from `git describe` in Makefile targets
- GoReleaser v2 config: cross-platform binaries, SHA256 checksums, CycloneDX SBOM, cosign signatures, Docker multi-arch images (GHCR), Homebrew tap formula

**CI/CD**
- Full production CI pipeline: code quality (gofmt, go vet, golangci-lint), cross-platform tests (ubuntu/macos/windows), 5-target build matrix, security scanning (govulncheck, gosec, SARIF), dependency review
- Release workflow: GoReleaser-powered, triggered on `v*` tags
- CodeQL analysis: weekly schedule + push/PR triggers
- Dependabot: Go modules and GitHub Actions auto-updates

**Security**
- Multi-stage distroless Dockerfile (nonroot user, no shell)
- `.dockerignore` minimizing build context
- cosign keyless signing in release pipeline
- SBOM (CycloneDX) generated per release artifact

**Developer Tooling**
- `.golangci.yml`: 20+ linters enabled including staticcheck, revive, gocritic, misspell
- `.editorconfig`: Go tabs, 2-space YAML/JSON enforced
- `.pre-commit-config.yaml`: go fmt, go vet, go build, go test, golangci-lint, conventional commits

**GitHub Automation**
- `CODEOWNERS`: maintainer review required for all core paths
- Release Drafter: automated, labeled release notes with semantic version resolution
- Issue template `config.yml`: blank issues disabled; directs to Discussions / Docs / Security Advisories
- Question issue template added

**Documentation**
- Professional README with badges, feature comparison table, installation matrix, command reference
- `docs/architecture.md`: 6 ADRs, output system design, CI/CD architecture diagram, security model, testing strategy
- `docs/installation.md`: all platforms, shell completions, checksum verification

**Tests**
- `internal/build` unit tests: 4 tests
- `internal/output` unit tests: 12 tests covering all rendering methods and edge cases
- `cmd/devopshub` integration tests: 18 tests covering root, version (plain/verbose/JSON), completion (all shells)
- `cmd/devopshub/cheat` integration tests: 14 tests covering structure, content, JSON validity, aliases, edge cases

---

[Unreleased]: https://github.com/devopshub-org/devopshub-cli/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/devopshub-org/devopshub-cli/releases/tag/v0.1.0
