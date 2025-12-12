# Makefile for math-skills project

# Binary name
BINARY_NAME=math-skills

# Build directory
BUILD_DIR=build

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build the project
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v

# Run the program
run:
	$(GOCMD) run .

# Run all tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -cover ./...
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests with race detector
test-race:
	@echo "Running tests with race detector..."
	$(GOTEST) -race -v ./...

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Format code
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

# Run linter (requires golangci-lint)
lint:
	@echo "Running linter..."
	golangci-lint run

# Install the binary
install: build
	@echo "Installing $(BINARY_NAME)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

# Show help
help:
	@echo "Available targets:"
	@echo "  make build         - Build the project"
	@echo "  make run          - Run the program"
	@echo "  make test         - Run all tests"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make test-race    - Run tests with race detector"
	@echo "  make bench        - Run benchmark tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make deps         - Download dependencies"
	@echo "  make fmt          - Format code"
	@echo "  make lint         - Run linter"
	@echo "  make install      - Install the binary"
	@echo "  make help         - Show this help message"

# Default target
.DEFAULT_GOAL := build

.PHONY: build run test test-coverage test-race bench clean deps fmt lint install help
