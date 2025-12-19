// Package main provides statistical calculations for math-skills program.
// This file contains unit tests for pure mathematical functions and
// integration tests for file I/O operations.
package main

import "testing"

// ============================================================================
// UNIT TESTS - Pure functions with no external dependencies
// ============================================================================

func TestAverage(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{"simple positive integers", []float64{1, 2, 3, 4, 5}, 3.0},
		{"single number", []float64{42}, 42.0},
		{"two numbers", []float64{10, 20}, 15.0},
		{"numbers with decimals", []float64{1.5, 2.5, 3.5}, 2.5},
		{"negative numbers", []float64{-10, -20, -30}, -20.0},
		{"mixed positive and negative", []float64{-5, 0, 5}, 0.0},
		{"large numbers", []float64{1000000, 2000000, 3000000}, 2000000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := average(tt.numbers)
			if result != tt.expected {
				t.Errorf("average(%v) = %f; expected %f", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{"odd count - simple", []float64{1, 3, 5}, 3.0},
		{"odd count - unsorted", []float64{5, 1, 3}, 3.0},
		{"even count - simple", []float64{1, 2, 3, 4}, 2.5},
		{"even count - unsorted", []float64{4, 1, 3, 2}, 2.5},
		{"single number", []float64{42}, 42.0},
		{"two numbers", []float64{10, 20}, 15.0},
		{"negative numbers", []float64{-10, -5, -20}, -10.0},
		{"mixed positive and negative", []float64{-5, 0, 5, 10}, 2.5},
		{"all same values", []float64{7, 7, 7, 7}, 7.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := median(tt.numbers)
			if result != tt.expected {
				t.Errorf("median(%v) = %f; expected %f", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestVariance(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{"simple case", []float64{4, 8, 6, 5, 3}, 2.96},
		{"all same values - zero variance", []float64{5, 5, 5, 5}, 0.0},
		{"single number", []float64{42}, 0.0},
		{"two numbers", []float64{1, 5}, 4.0},
		{"integers with integer variance", []float64{1, 2, 3, 4, 5}, 2.0},
		{"negative numbers", []float64{-10, -20, -30}, 66.66666666666667},
		{"mixed positive and negative", []float64{-5, 0, 5}, 16.666666666666668},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := variance(tt.numbers)
			delta := 0.00001
			if result < tt.expected-delta || result > tt.expected+delta {
				t.Errorf("variance(%v) = %f; expected %f", tt.numbers, result, tt.expected)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{"simple case", []float64{4, 8, 6, 5, 3}, 1.72046505340853},
		{"all same values - zero std dev", []float64{5, 5, 5, 5}, 0.0},
		{"single number", []float64{42}, 0.0},
		{"two numbers", []float64{1, 5}, 2.0},
		{"integers with integer std dev", []float64{1, 2, 3, 4, 5}, 1.4142135623730951},
		{"perfect square variance", []float64{1, 5, 9}, 3.265986323710904},
		{"negative numbers", []float64{-10, -20, -30}, 8.16496580927726},
		{"mixed positive and negative", []float64{-5, 0, 5}, 4.08248290463863},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := standardDeviation(tt.numbers)
			delta := 0.00001
			if result < tt.expected-delta || result > tt.expected+delta {
				t.Errorf("standardDeviation(%v) = %f; expected %f", tt.numbers, result, tt.expected)
			}
		})
	}
}

// ============================================================================
// INTEGRATION TESTS - Tests with filesystem dependencies
// ============================================================================

func TestReadNumbersFromFile(t *testing.T) {
	tests := []struct {
		name          string
		filename      string
		expected      []float64
		expectError   bool
		errorContains string
	}{
		{
			name:        "valid simple file",
			filename:    "testdata/valid_simple.txt",
			expected:    []float64{189, 113, 121, 114, 145, 110},
			expectError: false,
		},
		{
			name:        "single number",
			filename:    "testdata/single_number.txt",
			expected:    []float64{42},
			expectError: false,
		},
		{
			name:        "file with decimals",
			filename:    "testdata/with_decimals.txt",
			expected:    []float64{10.5, 20.3, 15.7},
			expectError: false,
		},
		{
			name:          "empty file",
			filename:      "testdata/empty.txt",
			expected:      nil,
			expectError:   true,
			errorContains: "no valid numbers",
		},
		{
			name:          "file does not exist",
			filename:      "testdata/nonexistent.txt",
			expected:      nil,
			expectError:   true,
			errorContains: "no such file",
		},
		{
			name:        "file with invalid data - skip invalid lines",
			filename:    "testdata/invalid_data.txt",
			expected:    []float64{189, 121, 145},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := readNumbersFromFile(tt.filename)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error containing '%s' but got none", tt.errorContains)
					return
				}
				if tt.errorContains != "" && !containsString(err.Error(), tt.errorContains) {
					t.Errorf("expected error containing '%s' but got: %v", tt.errorContains, err)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !floatSlicesEqual(result, tt.expected) {
				t.Errorf("readNumbersFromFile(%s) = %v; expected %v", tt.filename, result, tt.expected)
			}
		})
	}
}

// ============================================================================
// TEST HELPERS
// ============================================================================

// floatSlicesEqual compares two float64 slices for equality.
func floatSlicesEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// containsString checks if a string contains a substring.
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		func() bool {
			for i := 0; i <= len(s)-len(substr); i++ {
				if s[i:i+len(substr)] == substr {
					return true
				}
			}
			return false
		}())
}
