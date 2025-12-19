# Contributing to Math Skills

Thank you for your interest in contributing to Math Skills! This document provides guidelines for contributing to the project.

## 🎯 Project Goals

This project is designed to demonstrate:
- Test-Driven Development (TDD) principles
- Go best practices and idioms
- Statistical computation
- Clean, maintainable code

## 🚀 Getting Started

### Prerequisites

- Go 1.22.2 or higher
- Git
- A text editor or IDE

### Setting Up Development Environment

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR-USERNAME/math-skills.git
   cd math-skills
   ```
3. Run tests to verify setup:
   ```bash
   make test
   ```

## 📝 Development Workflow

### 1. Create a Branch

```bash
git checkout -b feature/your-feature-name
```

Use descriptive branch names:
- `feature/add-geometric-mean` for new features
- `fix/median-edge-case` for bug fixes
- `docs/update-readme` for documentation
- `test/improve-coverage` for test improvements

### 2. Follow TDD Principles

We use Test-Driven Development (TDD):

1. **🔴 RED**: Write a failing test
   ```bash
   go test -v -run TestYourFunction
   # Test should fail
   ```

2. **🟢 GREEN**: Write minimal code to pass the test
   ```bash
   go test -v -run TestYourFunction
   # Test should pass
   ```

3. **🔵 REFACTOR**: Improve the code
   ```bash
   go test -v
   # All tests should still pass
   ```

### 3. Writing Tests

Use **table-driven tests** (Go idiom):

```go
func TestYourFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    []float64
        expected float64
    }{
        {name: "case 1", input: []float64{1, 2, 3}, expected: 2.0},
        {name: "case 2", input: []float64{5, 5}, expected: 5.0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := yourFunction(tt.input)
            if result != tt.expected {
                t.Errorf("got %f, want %f", result, tt.expected)
            }
        })
    }
}
```

### 4. Code Style

#### Follow Go Conventions

- Run `gofmt` before committing:
  ```bash
  make fmt
  ```
- Use meaningful variable names
- Add comments for exported functions
- Keep functions small and focused

#### Documentation Comments

Every exported function must have a doc comment:

```go
// average calculates the arithmetic mean of a slice of numbers.
// It returns the sum of all numbers divided by the count.
//
// Parameters:
//   - numbers: slice of float64 values
//
// Returns:
//   - float64: the arithmetic mean
func average(numbers []float64) float64 {
    // implementation
}
```

### 5. Commit Messages

Use [Conventional Commits](https://www.conventionalcommits.org/):

**Format**: `<type>(<scope>): <subject>`

**Types**:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `test`: Adding or updating tests
- `refactor`: Code change that neither fixes a bug nor adds a feature
- `chore`: Maintenance tasks

**Examples**:
```bash
feat: add geometric mean calculation
fix: correct median calculation for even-length arrays
docs: update README with installation instructions
test: add edge cases for variance function
refactor: simplify average calculation logic
```

## ✅ Before Submitting

### Checklist

- [ ] All tests pass: `make test`
- [ ] Code is formatted: `make fmt`
- [ ] Added tests for new functionality
- [ ] Updated documentation if needed
- [ ] Commit messages follow conventions
- [ ] No unnecessary files committed

### Run Full Test Suite

```bash
# Run all tests
make test

# Check test coverage
make test-coverage

# Run benchmarks
make bench
```

## 🐛 Reporting Bugs

When reporting bugs, please include:

1. **Description**: What happened?
2. **Expected behavior**: What should have happened?
3. **Steps to reproduce**:
   ```bash
   go run . data.txt
   # Error occurs here
   ```
4. **Environment**:
   - Go version: `go version`
   - OS: Linux/macOS/Windows
5. **Sample data**: Attach the input file if relevant

## 💡 Suggesting Features

When suggesting features:

1. **Use case**: Why is this feature needed?
2. **Proposed solution**: How should it work?
3. **Alternatives**: What other approaches did you consider?
4. **Impact**: Will this be a breaking change?

## 📚 Resources

### Go Learning Resources

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go Testing Package](https://pkg.go.dev/testing)

### Statistical Resources

- [Khan Academy Statistics](https://www.khanacademy.org/math/statistics-probability)
- [Wikipedia: Variance](https://en.wikipedia.org/wiki/Variance)

## 🤝 Code of Conduct

- Be respectful and inclusive
- Welcome newcomers
- Focus on constructive feedback
- Help others learn

## ❓ Questions?

If you have questions:
1. Check existing documentation
2. Search closed issues
3. Open a new issue with the `question` label

Thank you for contributing! 🎉
