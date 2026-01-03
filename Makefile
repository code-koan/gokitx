# Makefile for Go project

# Go executable
GO ?= go

# Lint tool
GOLANGCI_LINT ?= golangci-lint

# Default target
.PHONY: all
all: fmt lint test

# Install dependencies
.PHONY: deps
deps:
	$(GO) mod tidy
	$(GO) mod download

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...
	@gofmt -s -w .

# Lint code
.PHONY: lint
lint:
	@echo "Running golangci-lint..."
	$(GOLANGCI_LINT) run ./...

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test ./... -v

# Run tests with coverage
.PHONY: coverage
coverage:
	@echo "Running tests with coverage..."
	$(GO) test ./... -coverprofile=coverage.out
	@echo "Coverage report:"
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

# Clean generated files
.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -f coverage.out coverage.html
	find . -type f -name '*.exe' -delete
	find . -type f -name '*.test' -delete
