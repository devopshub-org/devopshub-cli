<div align="center">

# DevOpsHub CLI

**The world's best open-source DevOps toolkit for engineers.**

[![CI](https://github.com/devopshub-org/devopshub-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/devopshub-org/devopshub-cli/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org/)
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
| K8s cluster inspection | ✅ `cluster inspect/events/pods/nodes` | Requires kubectl + plugins |
| Cloud credential helper | ✅ AWS, GCP, Azure authentication | Fragmented SDK tools |
| Config management | ✅ `~/.devopshub/config.yaml` | Manual file editing |
| Plugin system | ✅ Binary plugin discovery | Closed ecosystem |
| JSON output | ✅ `--json` on every command | Rarely consistent |
| No-color mode | ✅ `--no-color` | Often missing |
| Shell completions | ✅ bash/zsh/fish/powershell | Varies |
| Embedded cheatsheets | ✅ Docker, K8s, Linux | External references |
| Open source | ✅ Apache 2.0 | Varies |

---

## 🚀 Quickstart

```bash
# Install (Linux/macOS)
curl -sSfL https://raw.githubusercontent.com/devopshub-org/devopshub-cli/main/scripts/install.sh | bash

# Verify installation
devopshub version

# Inspect Kubernetes cluster health
devopshub cluster inspect

# Check active cloud provider authentication
devopshub cloud auth

# Initialize user configuration
devopshub config init

# Docker cheatsheet
devopshub cheat docker

# Structured JSON output for automation
devopshub cluster pods --json | jq '.[] | select(.Status != "Running")'
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
docker run --rm ghcr.io/devopshub-org/devopshub-cli:latest version
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

### `devopshub cluster`
Inspect and monitor Kubernetes clusters.

```bash
devopshub cluster inspect                # Health overview (Nodes, Pods, Warning Events)
devopshub cluster events                 # Recent cluster events (--type Warning --limit 50)
devopshub cluster pods                   # Pod status and resource requests (--namespace)
devopshub cluster nodes                  # Node capacity, conditions, and roles
```

### `devopshub cloud`
Cloud provider credential inspection and authentication helper.

```bash
devopshub cloud auth                     # Check AWS, GCP, and Azure credentials
devopshub cloud auth --provider aws      # Filter by specific cloud provider
```

### `devopshub config`
Manage user configuration settings (`~/.devopshub/config.yaml`).

```bash
devopshub config init                    # Create default configuration file
devopshub config view                    # View current active configuration
devopshub config get default_namespace   # Read specific configuration value
devopshub config set default_namespace app # Update configuration key value
```

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
│   ├── cloud/               # cloud subcommand package
│   ├── cluster/             # cluster subcommand package
│   ├── config/              # config subcommand package
│   ├── main.go              # Binary entry point
│   ├── root.go              # Root Cobra command
│   ├── version.go           # version command
│   └── completion.go        # shell completion command
├── internal/
│   ├── build/               # Build metadata (ldflags injection)
│   ├── config/              # User configuration loading (Viper)
│   └── output/              # Terminal output primitives
├── pkg/                     # Core library packages
│   ├── cloud/               # Cloud provider authentication providers
│   ├── k8s/                 # Kubernetes client wrappers and Fake inspect suites
│   └── plugin/              # Executable-based plugin discovery manager
├── docs/                    # Architecture + command documentation
│   └── adr/                 # Architecture Decision Records
├── scripts/                 # Install script, build helpers
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

---

## 🗺️ Roadmap

| Version | Status | Features |
|---|---|---|
| v0.1.0 | ✅ Completed | CLI foundation, version, cheat commands, full CI/CD |
| v0.2.0 | ✅ Completed | Kubernetes cluster inspection, cloud credential helpers, config, plugins |
| v0.3.0 | 📋 Planned | CI/CD pipeline helpers, infrastructure drift detection, TUI |
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
