# Installation Guide

## Prerequisites

- **Go 1.22+** (build from source only)
- **curl** (install script only)

---

## Homebrew (macOS / Linux)

```bash
brew install devopshub-org/tap/devopshub

# Verify
devopshub version
```

---

## Install Script (Linux / macOS)

```bash
# Latest version
curl -sSfL https://raw.githubusercontent.com/devopshub-org/devopshub-cli/main/scripts/install.sh | bash

# Specific version
curl -sSfL https://raw.githubusercontent.com/devopshub-org/devopshub-cli/main/scripts/install.sh | bash -s -- --version v0.1.0

# Custom install directory
curl -sSfL https://raw.githubusercontent.com/devopshub-org/devopshub-cli/main/scripts/install.sh | bash -s -- --dir ~/.local/bin
```

---

## Docker

```bash
# Run a command
docker run --rm ghcr.io/devopshub-org/devopshub-cli:latest cheat docker

# Pin to a specific version
docker run --rm ghcr.io/devopshub-org/devopshub-cli:v0.1.0 version

# Use as a base image
FROM ghcr.io/devopshub-org/devopshub-cli:v0.1.0
```

---

## Go Install

```bash
go install github.com/devopshub-org/devopshub-cli/cmd/devopshub@latest
```

> **Note:** `go install` builds without version metadata. Use the install script or Homebrew for production use.

---

## Build from Source

```bash
git clone https://github.com/devopshub-org/devopshub-cli.git
cd devopshub-cli

# Build for current platform
make build

# Install to GOPATH/bin
make install

# Cross-compile for all platforms
make build-all
```

---

## Shell Completions

### Bash
```bash
# Current session
source <(devopshub completion bash)

# Persist
devopshub completion bash > /etc/bash_completion.d/devopshub
```

### Zsh
```bash
source <(devopshub completion zsh)
# Persist: add to ~/.zshrc
```

### Fish
```bash
devopshub completion fish | source
devopshub completion fish > ~/.config/fish/completions/devopshub.fish
```

### PowerShell
```powershell
devopshub completion powershell | Out-String | Invoke-Expression
```

---

## Verifying a Release

Every release artifact has a corresponding SHA256 checksum in `checksums.txt`:

```bash
# Download binary
curl -sSfL -O https://github.com/devopshub-org/devopshub-cli/releases/download/v0.1.0/devopshub_0.1.0_linux_x86_64.tar.gz

# Download checksums
curl -sSfL -O https://github.com/devopshub-org/devopshub-cli/releases/download/v0.1.0/checksums.txt

# Verify
sha256sum --check --ignore-missing checksums.txt
```
