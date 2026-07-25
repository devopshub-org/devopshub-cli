# ADR 002: Executable-Based Plugin Architecture

## Status
Accepted

## Context
DevOpsHub CLI requires an extensible plugin system allowing community members and platform teams to extend `devopshub` with custom subcommands without needing to recompile the core CLI binary or modify the core codebase.

Key requirements:
1. Cross-language plugin development (Go, Python, Bash, Rust, Node.js).
2. Zero dynamic library dependency headaches (avoid C-GO / `plugin` package limitations in Go).
3. Seamless CLI integration (plugins behave like native subcommands).
4. Discoverability via `$PATH` and user configuration path (`~/.devopshub/plugins/`).

## Decision
We adopt the **Executable-Based Subcommand Plugin Pattern** (pioneered by `git` and `kubectl`).

### Naming & Discovery Scheme
- Any executable binary in `$PATH` or `~/.devopshub/plugins/` whose name begins with `devopshub-` is recognized as a plugin subcommand.
- Example: An executable binary named `devopshub-terraform` becomes accessible as `devopshub terraform`.

### Execution Model
- When `devopshub` receives a subcommand it does not natively implement, it scans for `devopshub-<subcommand>`.
- If found, `devopshub` invokes the executable binary, passing remaining command-line flags and arguments directly.
- Standard input (`stdin`), standard output (`stdout`), standard error (`stderr`), and environment variables (including `DEVOPSHUB_*` configuration options) are inherited by the plugin.

## Consequences
### Positive
- Extremely low barrier to entry: any executable or script can be a plugin.
- Completely decouples plugin lifecycle from `devopshub-cli` release cycle.
- Avoids Go `plugin` package platform constraints (which only work on Linux/macOS with matching compiler versions).

### Negative
- Inter-process communication (IPC) relies on CLI arguments/environment rather than rich in-process Go interfaces.
- Discovery requires searching directory paths on command invocation.
