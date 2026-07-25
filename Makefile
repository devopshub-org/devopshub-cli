.PHONY: all build clean test lint run help

BINARY_NAME=devopshub
MAIN_PATH=./cmd/devopshub

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/

test:
	@echo "Running tests..."
	go test -v ./...

lint:
	@echo "Running linter..."
	golangci-lint run ./... || go vet ./...

run:
	go run $(MAIN_PATH)

help:
	@echo "Usage:"
	@echo "  make build  - Build binary"
	@echo "  make clean  - Remove built binary"
	@echo "  make test   - Run tests"
	@echo "  make lint   - Run linters"
	@echo "  make run    - Run CLI"
