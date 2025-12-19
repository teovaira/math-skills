# Makefile for math-skills project
# A Makefile automates common development tasks

# ==============================================================================
# CONFIGURATION VARIABLES
# ==============================================================================

# Name of the compiled program
BINARY_NAME=math-skills

# Directory where compiled binaries will be stored
BUILD_DIR=build

# Go command shortcuts (so we can type $(GOBUILD) instead of "go build")
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

# ==============================================================================
# BUILD TARGETS
# ==============================================================================

# Default target - runs when you type just "make"
# Creates the executable binary in build/ directory
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v

# Run the program (note: you still need to provide filename as argument)
# Usage: make run ARGS="data.txt"
run:
	$(GOCMD) run .

# ==============================================================================
# TESTING TARGETS
# ==============================================================================

# Run all unit and integration tests
# -v = verbose (shows test names)
# ./... = test all packages recursively
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests and generate coverage report
# Creates coverage.html file you can open in a browser
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -cover ./...
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run performance benchmarks
# -bench=. runs all benchmark functions
# -benchmem shows memory allocations
bench:
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem

# ==============================================================================
# CODE QUALITY TARGETS
# ==============================================================================

# Automatically format all Go code according to Go standards
# This fixes indentation, spacing, etc.
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

# Run linter to check for code quality issues
# NOTE: Requires golangci-lint to be installed separately
# Install: https://golangci-lint.run/usage/install/
lint:
	@echo "Running linter..."
	golangci-lint run

# ==============================================================================
# CLEANUP TARGETS
# ==============================================================================

# Remove all generated files (binaries, coverage reports)
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# ==============================================================================
# UTILITY TARGETS
# ==============================================================================

# Tidy up Go module dependencies
# Removes unused dependencies and adds missing ones
deps:
	@echo "Tidying Go modules..."
	$(GOMOD) tidy

# Install the binary to your system PATH
# After this, you can run "math-skills" from anywhere
install: build
	@echo "Installing $(BINARY_NAME)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

# Show available make commands
help:
	@echo "Available targets:"
	@echo "  make build         - Build the project"
	@echo "  make run           - Run the program"
	@echo "  make test          - Run all tests"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make bench         - Run benchmark tests"
	@echo "  make fmt           - Format code"
	@echo "  make lint          - Run linter (requires golangci-lint)"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make deps          - Tidy Go modules"
	@echo "  make install       - Install the binary"
	@echo "  make help          - Show this help message"

# ==============================================================================
# SPECIAL MAKE DIRECTIVES
# ==============================================================================

# If you type just "make" with no target, it will run "build"
.DEFAULT_GOAL := build

# These are command names, not files
# Without this, if you had a file named "test", Make would get confused
.PHONY: build run test test-coverage bench fmt lint clean deps install help
