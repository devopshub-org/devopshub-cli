# devopshub-cli

The open-source DevOps toolkit for developers, platform engineers, and SREs.

## 🚀 Overview

`devopshub-cli` is a powerful, extensible command-line tool designed to streamline DevOps workflows, manage cloud infrastructure, inspect Kubernetes clusters, and automate routine operational tasks.

## 📂 Project Structure

```text
devopshub-cli/
├── .github/             # GitHub templates and CI workflows
├── cmd/                 # Application entrypoints (e.g. cmd/devopshub/main.go)
├── internal/            # Private application packages
├── pkg/                 # Public library code
├── docs/                # Project documentation
├── examples/            # Example configurations and command guides
├── assets/              # Logos, diagrams, and media
├── scripts/             # Build and utility scripts
├── tests/               # Integration and E2E test suites
├── README.md            # Project overview
├── LICENSE              # Apache 2.0 License
├── CONTRIBUTING.md      # Guidelines for contributing
├── ROADMAP.md           # Product roadmap and feature status
├── CHANGELOG.md         # Release history
├── SECURITY.md          # Vulnerability reporting policy
├── CODE_OF_CONDUCT.md   # Community guidelines
├── Makefile             # Task automation runner
└── go.mod               # Go module file
```

## 🛠️ Quick Start

### Prerequisites
- [Go](https://golang.org/dl/) 1.22 or higher
- `make` (optional)

### Installation

```bash
# Clone the repository
git clone https://github.com/devopshub-org/devopshub-cli.git
cd devopshub-cli

# Build binary
make build

# Run the CLI
./bin/devopshub --help
```

## 🤝 Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) to get started.

## 📄 License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.
