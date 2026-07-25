.PHONY: all build build-all clean test test-coverage lint fmt vet \
        install release-snapshot tidy help

# ─── Configuration ────────────────────────────────────────────────────────────

BINARY_NAME     := devopshub
MODULE          := github.com/devopshub-org/devopshub-cli
MAIN_PATH       := ./cmd/devopshub
BIN_DIR         := bin
DIST_DIR        := dist
COVERAGE_FILE   := coverage.out
COVERAGE_HTML   := coverage.html

VERSION         ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT          ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE            ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILT_BY        ?= make

GO              := go
GOFLAGS         ?=
LDFLAGS         := -s -w \
                   -X $(MODULE)/internal/build.Version=$(VERSION) \
                   -X $(MODULE)/internal/build.Commit=$(COMMIT) \
                   -X $(MODULE)/internal/build.Date=$(DATE) \
                   -X $(MODULE)/internal/build.BuiltBy=$(BUILT_BY)

# ─── Default ──────────────────────────────────────────────────────────────────

all: lint test build

# ─── Build ────────────────────────────────────────────────────────────────────

## build: Build binary for the current platform
build:
	@echo "→ Building $(BINARY_NAME) $(VERSION)..."
	@mkdir -p $(BIN_DIR)
	$(GO) build $(GOFLAGS) -trimpath -ldflags="$(LDFLAGS)" \
	  -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "✓ Binary: $(BIN_DIR)/$(BINARY_NAME)"

## build-all: Cross-compile for all supported platforms
build-all:
	@echo "→ Cross-compiling for all platforms..."
	@mkdir -p $(BIN_DIR)
	GOOS=linux   GOARCH=amd64  CGO_ENABLED=0 $(GO) build -trimpath -ldflags="$(LDFLAGS)" -o $(BIN_DIR)/$(BINARY_NAME)-linux-amd64   $(MAIN_PATH)
	GOOS=linux   GOARCH=arm64  CGO_ENABLED=0 $(GO) build -trimpath -ldflags="$(LDFLAGS)" -o $(BIN_DIR)/$(BINARY_NAME)-linux-arm64   $(MAIN_PATH)
	GOOS=darwin  GOARCH=amd64  CGO_ENABLED=0 $(GO) build -trimpath -ldflags="$(LDFLAGS)" -o $(BIN_DIR)/$(BINARY_NAME)-darwin-amd64  $(MAIN_PATH)
	GOOS=darwin  GOARCH=arm64  CGO_ENABLED=0 $(GO) build -trimpath -ldflags="$(LDFLAGS)" -o $(BIN_DIR)/$(BINARY_NAME)-darwin-arm64  $(MAIN_PATH)
	GOOS=windows GOARCH=amd64  CGO_ENABLED=0 $(GO) build -trimpath -ldflags="$(LDFLAGS)" -o $(BIN_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	@echo "✓ All binaries in $(BIN_DIR)/"

## install: Install binary to GOPATH/bin
install:
	@echo "→ Installing $(BINARY_NAME)..."
	$(GO) install $(GOFLAGS) -trimpath -ldflags="$(LDFLAGS)" $(MAIN_PATH)
	@echo "✓ Installed to $(shell go env GOPATH)/bin/$(BINARY_NAME)"

# ─── Tests ────────────────────────────────────────────────────────────────────

## test: Run unit tests
test:
	@echo "→ Running tests..."
	$(GO) test -v -race -count=1 ./...

## test-coverage: Run tests with coverage report
test-coverage:
	@echo "→ Running tests with coverage..."
	$(GO) test -race -count=1 -coverprofile=$(COVERAGE_FILE) -covermode=atomic ./...
	$(GO) tool cover -func=$(COVERAGE_FILE)
	$(GO) tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "✓ Coverage report: $(COVERAGE_HTML)"

# ─── Code Quality ─────────────────────────────────────────────────────────────

## lint: Run golangci-lint
lint:
	@echo "→ Running golangci-lint..."
	golangci-lint run ./...

## fmt: Format source code
fmt:
	@echo "→ Formatting code..."
	$(GO) fmt ./...
	goimports -w ./...

## vet: Run go vet
vet:
	@echo "→ Running go vet..."
	$(GO) vet ./...

## tidy: Tidy Go modules
tidy:
	@echo "→ Tidying modules..."
	$(GO) mod tidy
	$(GO) mod verify

# ─── Release ──────────────────────────────────────────────────────────────────

## release-snapshot: Build a local release snapshot with GoReleaser
release-snapshot:
	@echo "→ Building release snapshot..."
	goreleaser release --snapshot --clean

# ─── Cleanup ──────────────────────────────────────────────────────────────────

## clean: Remove build artifacts
clean:
	@echo "→ Cleaning..."
	@rm -rf $(BIN_DIR) $(DIST_DIR) $(COVERAGE_FILE) $(COVERAGE_HTML)
	@echo "✓ Clean"

# ─── Help ─────────────────────────────────────────────────────────────────────

## help: Show this help message
help:
	@echo "DevOpsHub CLI — Makefile"
	@echo ""
	@echo "Usage: make <target>"
	@echo ""
	@grep -E '^## [a-zA-Z_-]+:' $(MAKEFILE_LIST) | \
	  sort | \
	  awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' | \
	  sed 's/## //'
