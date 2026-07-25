# DevOpsHub CLI — Roadmap

This document outlines the vision, milestones, and feature status for `devopshub-cli`.

## 🎯 Vision

To become the GitHub standard DevOps CLI — the first tool every DevOps engineer installs.

**Goal:** 100,000 GitHub Stars · CNCF Ecosystem · Enterprise Adoption

---

## 🗺️ Release Milestones

### ✅ v0.1.0 — Foundation (Milestone 1)
**Status:** Completed

The architectural foundation. A compilable, testable, releasable CLI.

- [x] Cobra CLI framework
- [x] `version` command (plain / verbose / JSON)
- [x] `completion` command (bash/zsh/fish/powershell)
- [x] `cheat docker` — Docker cheatsheet (50+ commands)
- [x] `cheat k8s` — Kubernetes cheatsheet (60+ commands)
- [x] `cheat linux` — Linux cheatsheet (100+ commands)
- [x] `--no-color` / `--json` / `--debug` / `--quiet` global flags
- [x] Internal `output` package (consistent terminal UX)
- [x] Internal `build` package (ldflags version injection)
- [x] Production CI/CD (quality · test · build · security · release)
- [x] GoReleaser (cross-platform binaries, checksums, SBOM, Docker)
- [x] Distroless Docker image (GHCR)
- [x] Homebrew tap formula
- [x] CodeQL + govulncheck + gosec security scanning
- [x] golangci-lint (20+ linters)
- [x] Pre-commit hooks + EditorConfig
- [x] CODEOWNERS + Release Drafter + Dependabot
- [x] Full test suite (90%+ coverage target)
- [x] Architecture docs + ADRs

---

### ✅ v0.2.0 — Kubernetes & Cloud (Milestone 2)
**Status:** Completed

- [x] `cluster inspect` — Kubernetes cluster health overview
- [x] `cluster events` — Recent cluster events with filtering
- [x] `cluster pods` — Pod status with resource usage
- [x] `cluster nodes` — Node capacity and conditions
- [x] `cloud auth` — Cloud provider credential helper (AWS/GCP/Azure)
- [x] `config` command — User configuration file (`~/.devopshub/config.yaml`)
- [x] Viper integration for config file support
- [x] Plugin system foundation (ADR 002)

---

### 📋 v0.3.0 — Automation & CI/CD (Milestone 3)
**Status:** Planned

- [ ] `ci validate` — Validate GitHub Actions / GitLab CI syntax
- [ ] `drift detect` — Infrastructure drift detection
- [ ] `tui` — Interactive terminal UI (TUI) for cluster management
- [ ] `env diff` — Compare environment variable sets
- [ ] `secret scan` — Basic secrets detection in files/environment
- [ ] Shell plugin system (load custom commands)

---

### 🎯 v1.0.0 — Stable Release
**Status:** Long-term goal

- [ ] Stable public API (committed backward compatibility)
- [ ] Plugin marketplace
- [ ] VSCode extension
- [ ] Desktop application
- [ ] AI assistant integration
- [ ] CNCF project submission
- [ ] Enterprise support tiers

---

## 🌱 Community Goals

| Goal | Target |
|---|---|
| GitHub Stars | 100,000 |
| CNCF Sandbox | v1.0.0 |
| Package Managers | Homebrew · Scoop · APT · Nix |
| Integrations | VSCode · JetBrains · GitHub Actions Marketplace |

---

> Features and timeline are subject to change based on community feedback.
> Open a [GitHub Discussion](https://github.com/devopshub-org/devopshub-cli/discussions) to propose new features.
