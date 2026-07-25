# Architecture

This document describes the architecture and design decisions behind DevOpsHub CLI.

## Overview

DevOpsHub CLI is a single static binary built in Go. It uses the Cobra command framework and follows the standard Go project layout.

## Package Map

```
cmd/devopshub/          → Binary entry point + top-level commands
  cheat/                → Cheatsheet command package

internal/build/         → Version metadata (ldflags injection)
internal/output/        → Terminal output primitives (single source of truth)

pkg/                    → Public library packages (future — none in v0.1.0)
```

## Architecture Decision Records

### ADR-001 — Command Framework: Cobra

**Status:** Accepted  
**Decision:** Use `github.com/spf13/cobra`.  
**Rationale:** Industry standard for production Go CLIs (kubectl, helm, gh, k9s). Provides auto-generated completions, structured subcommand hierarchy, and the most mature plugin model in the Go ecosystem.  
**Alternatives Considered:** `urfave/cli` — smaller API but weaker ecosystem; rejected.

---

### ADR-002 — Logging: zerolog

**Status:** Accepted  
**Decision:** Use `github.com/rs/zerolog` for structured debug output.  
**Rationale:** Zero-allocation logger, fastest benchmark in the Go ecosystem, designed for CLI use (no goroutines, no global mutex on the hot path). Used only for `--debug` output.  
**Alternatives Considered:** `slog` (stdlib) — adequate but zerolog's ergonomics are superior for CLI debug output.

---

### ADR-003 — Output: Custom `internal/output` Package

**Status:** Accepted  
**Decision:** All user-facing text rendering goes through `internal/output.Printer`. Commands MUST NOT call `fmt.Println` directly.  
**Rationale:** Single point of control for `--no-color`, `--json`, `--quiet` flags. Consistent visual language across all commands. Testable via injected `io.Writer`.

---

### ADR-004 — Version Injection: ldflags

**Status:** Accepted  
**Decision:** Inject `Version`, `Commit`, `Date`, `BuiltBy` via `-ldflags` at build time.  
**Rationale:** Zero runtime cost. Standard pattern (GoReleaser-compatible). No config file required.

---

### ADR-005 — Configuration: No Viper in v0.1.0

**Status:** Accepted  
**Decision:** No Viper. Plain Cobra persistent flags only.  
**Rationale:** No persistent config file needed at this milestone. YAGNI. Viper will be added in v0.2.0 when a config file (`~/.devopshub/config.yaml`) is required for cloud provider credentials.

---

### ADR-006 — Cheatsheet Data: Embedded Structs

**Status:** Accepted  
**Decision:** Cheatsheet data is declared as typed Go structs in each command file. No external files, no `embed.FS`.  
**Rationale:** The data is small (< 200 entries per sheet). Struct embedding allows Go's type system to enforce JSON field names via struct tags. No I/O, no error handling, no deployment complexity.  
**Future:** If cheatsheets grow beyond ~500 entries, migrate to `embed.FS` with YAML/JSON source files.

---

## Output System

The `internal/output` package is the single source of truth for all terminal rendering.

### Design

```
output.Printer (struct)
  ├── out      io.Writer  ← injected; defaults to os.Stdout
  ├── errOut   io.Writer  ← injected; defaults to os.Stderr
  ├── noColor  bool
  ├── jsonMode bool
  └── quiet    bool

Methods:
  Success(msg)       → ✓ green  → stdout
  Error(msg)         → ✗ red    → stderr (always, even in quiet mode)
  Info(msg)          → ℹ blue   → stdout
  Warning(msg)       → ⚠ yellow → stdout
  Header(title)      →          → stdout
  Table(headers, rows)          → stdout
  Println/Printf                → stdout
  JSON(v)            → JSON     → stdout
  JSONMode() bool               → query current mode
```

### Flag Propagation

The root command's `PersistentPreRunE` rebuilds the printer from parsed flag values before every subcommand runs. Commands receive the printer via constructor injection.

---

## CI/CD Architecture

```
push / PR
  ├── quality    → gofmt · go vet · golangci-lint
  ├── test       → ubuntu · macos · windows (race detector + coverage)
  ├── build      → 5-platform cross-compilation matrix
  ├── security   → govulncheck · gosec · SARIF upload
  └── dependency-review  (PRs only)

push v* tag
  └── release → GoReleaser
        ├── binaries (linux/darwin/windows × amd64/arm64)
        ├── archives (tar.gz / zip)
        ├── checksums.txt (SHA256)
        ├── SBOM (CycloneDX)
        ├── cosign signatures
        ├── Docker images → GHCR (multi-arch manifest)
        ├── Homebrew formula → devopshub-org/homebrew-tap
        └── GitHub Release
```

---

## Security Model

- No network access at runtime (cheatsheets are embedded).
- No file I/O at runtime (no config file in v0.1.0).
- Binary is built with `-trimpath` (no local paths in binary).
- Binary is built with `-s -w` (stripped symbols + DWARF).
- Docker image uses distroless/static with `nonroot` user.
- Releases are signed with `cosign` (keyless, OIDC-based).
- SBOM (CycloneDX) generated per release.

---

## Testing Strategy

| Layer | Type | Target Coverage |
|---|---|---|
| `internal/build` | Unit | 100% |
| `internal/output` | Unit | 90%+ |
| `cmd/devopshub` (root, version, completion) | Integration | 85%+ |
| `cmd/devopshub/cheat` | Integration | 85%+ |

All tests use injected `io.Writer` — no global state, no `os.Stdout` capture.
