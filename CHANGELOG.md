# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-12-11

### Added
- Initial release of Math Skills
- **Statistical Functions**:
  - `average()`: Calculate arithmetic mean
  - `median()`: Calculate median value
  - `variance()`: Calculate variance
  - `standardDeviation()`: Calculate standard deviation
- **File I/O**:
  - `readNumbersFromFile()`: Read numbers from file (one per line)
  - Support for decimal numbers
  - Graceful handling of invalid data (skip non-numeric lines)
  - Error handling for missing files and empty files
- **Main Program**:
  - Command-line interface
  - Proper error messages
  - Results rounded to nearest integer
- **Testing**:
  - 37 comprehensive tests (31 unit + 6 file reading)
  - Table-driven test pattern
  - Test coverage for edge cases
  - Test data files in `testdata/` directory
- **Documentation**:
  - Comprehensive README.md
  - Contributing guidelines (CONTRIBUTING.md)
  - Makefile for common tasks
  - Detailed function documentation
  - Statistical formula explanations
- **Build Tools**:
  - Makefile with targets: build, test, clean, fmt
  - Go module support
  - `.gitignore` for Go projects

### Features
- ✅ Handles both positive and negative numbers
- ✅ Supports decimal values
- ✅ Skips invalid lines in input files
- ✅ Provides clear error messages
- ✅ All results rounded to integers as per specification

### Test Coverage
- Average: 7 test cases
- Median: 9 test cases (odd/even counts, sorted/unsorted)
- Variance: 7 test cases
- Standard Deviation: 8 test cases
- File Reading: 6 test cases (valid/invalid/missing files)

---

## Release Notes

### Version 1.0.0

This is the first stable release of Math Skills, a Go program for calculating statistical measures from data files.

**Highlights**:
- Complete implementation of 4 statistical functions
- Comprehensive test suite with 100% passing tests
- Professional documentation and build tools
- Follows Go best practices and TDD methodology

**Known Limitations**:
- Only supports newline-separated values (one number per line)
- Results are rounded to integers (as per project specification)
- No support for weighted statistics

**Future Enhancements** (Potential):
- Add geometric mean and harmonic mean
- Support for CSV file format
- Streaming mode for large files
- Performance benchmarks

---

## How to Read This Changelog

### Categories

- **Added**: New features
- **Changed**: Changes in existing functionality
- **Deprecated**: Soon-to-be removed features
- **Removed**: Removed features
- **Fixed**: Bug fixes
- **Security**: Security fixes

### Version Format

Version numbers follow [Semantic Versioning](https://semver.org/):
- **MAJOR**: Incompatible API changes
- **MINOR**: Backwards-compatible functionality additions
- **PATCH**: Backwards-compatible bug fixes

---

[1.0.0]: https://github.com/yourusername/math-skills/releases/tag/v1.0.0
