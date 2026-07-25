# Contributing to devopshub-cli

First off, thank you for considering contributing to `devopshub-cli`! It's people like you that make devopshub-cli a great tool for the community.

## How Can I Contribute?

### Reporting Bugs
Before creating a bug report, please check existing issues to avoid duplicates. When creating a bug report, include:
- A clear descriptive title
- Steps to reproduce the problem
- Expected vs. actual behavior
- OS and Go version info

### Requesting Features
Feature requests are tracked as GitHub issues. When suggesting a feature, please include:
- A clear, descriptive title
- Explanation of the use case and why it would be beneficial
- Any mockups or design ideas

### Pull Requests
1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/amazing-feature`).
3. Commit your changes (`git commit -m 'feat: add amazing feature'`).
4. Push to the branch (`git push origin feature/amazing-feature`).
5. Open a Pull Request.

## Development Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/devopshub-cli.git
cd devopshub-cli

# Build locally
make build

# Run unit tests
make test
```

## Code Style

- Follow standard Go formatting guidelines (`gofmt` / `go vet`).
- Keep code clean, readable, and well-documented.
- Write tests for new functionality.

## Code of Conduct

Please note that this project is released with a Contributor Code of Conduct. By participating, you agree to abide by its terms. See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).
