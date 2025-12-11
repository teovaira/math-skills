# Math Skills

A Go program that calculates statistical measures (average, median, variance, and standard deviation) from a data file.

## 📋 Project Description

This program reads numerical data from a file (one number per line) and calculates four fundamental statistical measures:

- **Average (Mean)**: The sum of all values divided by the count
- **Median**: The middle value when numbers are sorted
- **Variance**: A measure of how spread out the numbers are from the average
- **Standard Deviation**: The square root of variance, representing typical distance from average

All results are rounded to the nearest integer.

## 🚀 Installation

### Prerequisites

- Go 1.22.2 or higher

### Clone and Build

```bash
git clone <your-repo-url>
cd math-skills
go build -o math-skills
```

## 💻 Usage

### Running with `go run`

```bash
go run . <filename>
```

### Running the compiled binary

```bash
./math-skills <filename>
```

### Example

Given a file `data.txt`:
```
189
113
121
114
145
110
```

Run:
```bash
go run . data.txt
```

Output:
```
Average: 132
Median: 118
Variance: 785
Standard Deviation: 28
```

## 📁 Input File Format

- One number per line
- Supports both integers and decimals
- Empty lines are ignored
- Invalid lines (non-numeric) are skipped silently

## 🧪 Testing

Run all tests:
```bash
go test -v
```

Run specific test:
```bash
go test -v -run TestAverage
```

Run tests with coverage:
```bash
go test -cover
```

## 📊 Test Coverage

The project includes comprehensive test coverage:
- **31 unit tests** for statistical functions
- **6 integration tests** for file reading
- **Edge cases**: empty files, single numbers, invalid data, negative numbers

## 🏗️ Project Structure

```
math-skills/
├── main.go           # Entry point and main logic
├── stats.go          # Statistical calculation functions
├── stats_test.go     # Comprehensive unit tests
├── testdata/         # Test data files
│   ├── valid_simple.txt
│   ├── single_number.txt
│   ├── with_decimals.txt
│   ├── empty.txt
│   └── invalid_data.txt
├── go.mod            # Go module definition
├── README.md         # This file
└── Makefile          # Build and test automation
```

## 🔧 Development

### Running tests during development

```bash
make test
```

### Building the binary

```bash
make build
```

### Cleaning build artifacts

```bash
make clean
```

## 📚 Statistical Formulas

### Average
```
average = (x₁ + x₂ + ... + xₙ) / n
```

### Median
- If n is odd: middle element of sorted array
- If n is even: average of two middle elements

### Variance
```
variance = Σ(xᵢ - μ)² / n
```
where μ is the average

### Standard Deviation
```
standardDeviation = √variance
```

## 🎓 Learning Goals

This project demonstrates:
- ✅ Test-Driven Development (TDD)
- ✅ Table-driven tests
- ✅ File I/O operations
- ✅ Error handling
- ✅ Statistical calculations
- ✅ Go best practices and idioms

## 🤝 Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute to this project.

## 📝 License

This project is created for educational purposes as part of the Zone01 Athens curriculum.

## 👤 Author

Developed by a Zone01 Athens student learning Go programming and software development best practices.
