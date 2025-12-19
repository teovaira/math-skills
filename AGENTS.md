# Agent Instructions

## Project Overview

Go program that calculates statistical measures (average, median, variance, standard deviation) from a data file. Built with Go 1.22.2+ using only the standard library.

## Tech Stack

- **Language**: Go 1.22.2+
- **Dependencies**: None (standard library only)
- **Testing**: Go's built-in `testing` package
- **Build**: Make + Go toolchain

## Setup Commands

```bash
# Build binary
make build

# Run program
go run . data.txt
```

## Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make bench
```

## Code Style

### Go Conventions
- Use `gofmt` formatting (run `make fmt` before commits)
- Table-driven tests for all test functions
- Detailed function documentation with parameters and return values
- Error handling: explicit returns, contextual error messages

### Example Test Pattern
```go
func TestFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    []float64
        expected float64
    }{
        {"case description", []float64{1, 2, 3}, 2.0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := function(tt.input)
            if result != tt.expected {
                t.Errorf("got %f, want %f", result, tt.expected)
            }
        })
    }
}
```

## Project Structure

```
math-skills/
├── main.go              # Entry point (argument parsing, orchestration)
├── stats.go             # Statistical functions (average, median, variance, stddev)
├── io.go                # File I/O operations
├── stats_test.go        # Unit and integration tests
├── stats_bench_test.go  # Performance benchmarks
├── testdata/            # Test data files
└── Makefile             # Build automation
```

## File Organization Rules

- **main.go**: Command-line interface only, no business logic
- **stats.go**: Pure functions for calculations, detailed comments with formulas
- **io.go**: File I/O operations separate from calculations
- **Separation of concerns**: I/O, calculation, and presentation logic in separate files

## Agent Boundaries

### ✅ Always Do
- Run `make fmt` before any code changes
- Add table-driven tests for new functions
- Include statistical formulas in function comments
- Handle edge cases: empty input, single value, negative numbers
- Use `float64` for all numeric operations

### ⚠️ Ask First
- Adding external dependencies (project uses standard library only)
- Changing statistical formulas or rounding behavior
- Modifying public function signatures
- Adding new statistical measures

### 🚫 Never Do
- Skip tests when adding functions
- Use external math libraries (use `math` from stdlib)
- Modify test data in `testdata/` without documenting why
- Round intermediate calculations (only round final output)

## Build & Quality Checks

```bash
# Format code
make fmt

# Run linter (requires golangci-lint installed)
make lint

# Clean build artifacts
make clean

# Full validation workflow
make fmt && make test && make build
```

## Development Workflow

1. Write failing test (TDD - Red)
2. Implement minimal code to pass (TDD - Green)
3. Refactor for clarity (TDD - Refactor)
4. Run `make fmt && make test`
5. Commit with conventional commit message

## Statistical Implementation Notes

### Formulas Used
- **Average**: `sum / count`
- **Median**: Middle value (odd) or average of two middle values (even)
- **Variance**: `Σ(xᵢ - μ)² / n` where μ is average
- **Standard Deviation**: `√variance`

### Precision Rules
- Internal calculations use `float64`
- Final output rounded to nearest integer with `math.Round()`
- Test comparisons use delta of `0.00001` for float precision

## Common Commands Reference

| Command | Purpose |
|---------|---------|
| `make build` | Build executable binary |
| `make test` | Run all unit tests |
| `make test-coverage` | Generate HTML coverage report |
| `make bench` | Run performance benchmarks |
| `make fmt` | Format all Go files |
| `make clean` | Remove build artifacts |
| `go run . data.txt` | Run with input file |
