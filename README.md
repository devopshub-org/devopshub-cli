<div align="center">

# DevOpsHub CLI

**The world's best open-source DevOps toolkit for engineers.**

[![CI](https://github.com/devopshub-org/devopshub-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/devopshub-org/devopshub-cli/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/devopshub-org/devopshub-cli)](https://github.com/devopshub-org/devopshub-cli/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/devopshub-org/devopshub-cli)](https://goreportcard.com/report/github.com/devopshub-org/devopshub-cli)
[![codecov](https://codecov.io/gh/devopshub-org/devopshub-cli/branch/main/graph/badge.svg)](https://codecov.io/gh/devopshub-org/devopshub-cli)
[![CNCF Ecosystem](https://img.shields.io/badge/CNCF-Ecosystem-blue)](https://cncf.io/)
[![OpenSSF Best Practices](https://bestpractices.coreinfrastructure.org/projects/0/badge)](https://bestpractices.coreinfrastructure.org/)

A powerful, extensible CLI designed to streamline DevOps workflows, manage cloud infrastructure, inspect Kubernetes clusters, and automate operational tasks — all from one tool.

[**Quickstart**](#-quickstart) • [**Commands**](#-commands) • [**Installation**](#-installation) • [**Contributing**](#-contributing) • [**Roadmap**](ROADMAP.md)

</div>

---

## ✨ Why DevOpsHub?

| Feature | DevOpsHub | Other Tools |
|---|---|---|
| Single binary | ✅ | Varies |
| No runtime deps | ✅ | Often requires runtime |
| Cross-platform | ✅ Linux · macOS · Windows | Often Linux-only |
| JSON output | ✅ `--json` on every command | Rarely consistent |
| No-color mode | ✅ `--no-color` | Often missing |
| Shell completions | ✅ bash/zsh/fish/powershell | Varies |
| Docker/K8s/Linux cheatsheets | ✅ Embedded in binary | External references |
| Open source | ✅ Apache 2.0 | Varies |

---

## 🚀 Quickstart

```bash
# Install (Linux/macOS)
curl -sSfL https://raw.githubusercontent.com/devopshub-org/devopshub-cli/main/scripts/install.sh | bash

# Verify installation
devopshub version

# Docker cheatsheet
devopshub cheat docker

# Kubernetes cheatsheet
devopshub cheat k8s

# Linux cheatsheet
devopshub cheat linux

# JSON output for scripting
devopshub cheat docker --json | jq '.[] | .title'
```

---

## 📦 Installation

### Homebrew (macOS / Linux)
```bash
brew install devopshub-org/tap/devopshub
```

### Curl installer (Linux / macOS)
```bash
curl -sSfL https://raw.githubusercontent.com/devopshub-org/devopshub-cli/main/scripts/install.sh | bash
```

### Docker
```bash
docker run --rm ghcr.io/devopshub-org/devopshub-cli:latest cheat docker
```

### Build from source
```bash
git clone https://github.com/devopshub-org/devopshub-cli.git
cd devopshub-cli
make install
```

### Go install
```bash
go install github.com/devopshub-org/devopshub-cli/cmd/devopshub@latest
```

> See [docs/installation.md](docs/installation.md) for all platforms and package managers.

---

## 🛠️ Commands

### `devopshub version`
Print version and build metadata.

```bash
devopshub version                # Short version line
devopshub version --verbose      # Full build info table
devopshub version --json         # Machine-readable JSON
```

### `devopshub cheat`
Embedded DevOps cheatsheets. No internet required.

```bash
devopshub cheat docker           # Docker commands (50+ commands, 7 sections)
devopshub cheat k8s              # kubectl commands (60+ commands, 9 sections)
devopshub cheat linux            # Linux commands (100+ commands, 8 sections)
devopshub cheat docker --json    # Structured JSON output
```

### `devopshub completion`
Generate shell completion scripts.

```bash
devopshub completion bash        # Bash completions
devopshub completion zsh         # Zsh completions
devopshub completion fish        # Fish completions
devopshub completion powershell  # PowerShell completions
```

### Global Flags
All commands support:

| Flag | Short | Description |
|---|---|---|
| `--no-color` | | Disable ANSI color output |
| `--json` | `-j` | Output as structured JSON |
| `--debug` | | Enable debug logging |
| `--quiet` | `-q` | Suppress informational output |

---

## 📂 Project Structure

```
devopshub-cli/
├── cmd/devopshub/           # CLI entry point + command definitions
│   ├── cheat/               # cheat subcommand package
│   ├── main.go              # Binary entry point
│   ├── root.go              # Root Cobra command
│   ├── version.go           # version command
│   └── completion.go        # shell completion command
├── internal/
│   ├── build/               # Build metadata (ldflags injection)
│   └── output/              # Terminal output primitives
├── pkg/                     # Public library code (future)
├── docs/                    # Architecture + command documentation
├── examples/                # Example usage and configurations
├── assets/                  # Logos, diagrams, media
├── scripts/                 # Install script, build helpers
├── tests/                   # Integration + E2E tests
└── .github/                 # CI/CD, issue templates, automation
```

---

## 🤝 Contributing

We welcome contributions of all kinds! Read [CONTRIBUTING.md](CONTRIBUTING.md) to get started.

### Development setup
```bash
git clone https://github.com/devopshub-org/devopshub-cli.git
cd devopshub-cli

# Install pre-commit hooks
pip install pre-commit && pre-commit install

# Build
make build

# Test
make test

# Lint
make lint

# Coverage report
make test-coverage
```

### Conventional Commits
This project follows [Conventional Commits](https://www.conventionalcommits.org/):
```
feat: add new command
fix: resolve panic in version command
docs: update installation guide
chore(deps): bump cobra to v1.9.0
```

---

## 🗺️ Roadmap

| Version | Status | Features |
|---|---|---|
| v0.1.0 | 🚧 In Progress | CLI foundation, version, cheat commands, full CI/CD |
| v0.2.0 | 📋 Planned | Kubernetes cluster inspection, cloud credential helpers |
| v0.3.0 | 📋 Planned | CI/CD pipeline helpers, infrastructure drift detection |
| v1.0.0 | 🎯 Goal | Stable API, plugin marketplace, VSCode extension |

See [ROADMAP.md](ROADMAP.md) for full details.

---

## 🔒 Security

Please report security vulnerabilities via [GitHub Security Advisories](https://github.com/devopshub-org/devopshub-cli/security/advisories/new).

See [SECURITY.md](SECURITY.md) for our security policy.

---

## 📄 License

Licensed under the [Apache License 2.0](LICENSE).

---

<div align="center">
  <strong>Built with ❤️ for the DevOps community</strong><br/>
  <a href="https://github.com/devopshub-org/devopshub-cli/stargazers">⭐ Star us on GitHub</a>
</div>
