package main

import "testing"

// TestAverage tests the average calculation function with various inputs
func TestAverage(t *testing.T) {
	// Table-driven tests: each test case has input and expected output
	tests := []struct {
		name     string      // Description of the test case
		numbers  []float64   // Input data
		expected float64     // Expected result
	}{
		{
			name:     "simple positive integers",
			numbers:  []float64{1, 2, 3, 4, 5},
			expected: 3.0,
		},
		{
			name:     "single number",
			numbers:  []float64{42},
			expected: 42.0,
		},
		{
			name:     "two numbers",
			numbers:  []float64{10, 20},
			expected: 15.0,
		},
		{
			name:     "numbers with decimals",
			numbers:  []float64{1.5, 2.5, 3.5},
			expected: 2.5,
		},
		{
			name:     "negative numbers",
			numbers:  []float64{-10, -20, -30},
			expected: -20.0,
		},
		{
			name:     "mixed positive and negative",
			numbers:  []float64{-5, 0, 5},
			expected: 0.0,
		},
		{
			name:     "large numbers",
			numbers:  []float64{1000000, 2000000, 3000000},
			expected: 2000000.0,
		},
	}

	// Run each test case
	for _, tt := range tests {
		// t.Run creates a subtest for each case
		t.Run(tt.name, func(t *testing.T) {
			result := average(tt.numbers)
			if result != tt.expected {
				t.Errorf("average(%v) = %f; expected %f", tt.numbers, result, tt.expected)
			}
		})
	}
}

// TestMedian tests the median calculation function with various inputs
func TestMedian(t *testing.T) {
	// Table-driven tests for median
	tests := []struct {
		name     string      // Description of the test case
		numbers  []float64   // Input data
		expected float64     // Expected result
	}{
		{
			name:     "odd count - simple",
			numbers:  []float64{1, 3, 5},
			expected: 3.0,
		},
		{
			name:     "odd count - unsorted",
			numbers:  []float64{5, 1, 3},
			expected: 3.0,
		},
		{
			name:     "even count - simple",
			numbers:  []float64{1, 2, 3, 4},
			expected: 2.5, // (2 + 3) / 2
		},
		{
			name:     "even count - unsorted",
			numbers:  []float64{4, 1, 3, 2},
			expected: 2.5,
		},
		{
			name:     "single number",
			numbers:  []float64{42},
			expected: 42.0,
		},
		{
			name:     "two numbers",
			numbers:  []float64{10, 20},
			expected: 15.0,
		},
		{
			name:     "negative numbers",
			numbers:  []float64{-10, -5, -20},
			expected: -10.0,
		},
		{
			name:     "mixed positive and negative",
			numbers:  []float64{-5, 0, 5, 10},
			expected: 2.5, // (0 + 5) / 2
		},
		{
			name:     "all same values",
			numbers:  []float64{7, 7, 7, 7},
			expected: 7.0,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := median(tt.numbers)
			if result != tt.expected {
				t.Errorf("median(%v) = %f; expected %f", tt.numbers, result, tt.expected)
			}
		})
	}
}

// TestVariance tests the variance calculation function with various inputs
func TestVariance(t *testing.T) {
	// Table-driven tests for variance
	tests := []struct {
		name     string      // Description of the test case
		numbers  []float64   // Input data
		expected float64     // Expected result
	}{
		{
			name:     "simple case",
			numbers:  []float64{4, 8, 6, 5, 3},
			expected: 2.96, // Manual calculation from theory
		},
		{
			name:     "all same values - zero variance",
			numbers:  []float64{5, 5, 5, 5},
			expected: 0.0, // No spread = zero variance
		},
		{
			name:     "single number",
			numbers:  []float64{42},
			expected: 0.0, // Single value has no variance
		},
		{
			name:     "two numbers",
			numbers:  []float64{1, 5},
			expected: 4.0, // avg=3, (1-3)²=4, (5-3)²=4, sum=8, 8/2=4
		},
		{
			name:     "integers with integer variance",
			numbers:  []float64{1, 2, 3, 4, 5},
			expected: 2.0, // avg=3, sum of squares=10, 10/5=2
		},
		{
			name:     "negative numbers",
			numbers:  []float64{-10, -20, -30},
			expected: 66.66666666666667, // avg=-20
		},
		{
			name:     "mixed positive and negative",
			numbers:  []float64{-5, 0, 5},
			expected: 16.666666666666668, // avg=0
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := variance(tt.numbers)
			// Use a small delta for float comparison due to precision
			delta := 0.00001
			if result < tt.expected-delta || result > tt.expected+delta {
				t.Errorf("variance(%v) = %f; expected %f", tt.numbers, result, tt.expected)
			}
		})
	}
}

// TestStandardDeviation tests the standard deviation calculation function with various inputs
func TestStandardDeviation(t *testing.T) {
	// Table-driven tests for standard deviation
	tests := []struct {
		name     string      // Description of the test case
		numbers  []float64   // Input data
		expected float64     // Expected result (sqrt of variance)
	}{
		{
			name:     "simple case",
			numbers:  []float64{4, 8, 6, 5, 3},
			expected: 1.72046505340853, // √2.96
		},
		{
			name:     "all same values - zero std dev",
			numbers:  []float64{5, 5, 5, 5},
			expected: 0.0, // √0 = 0
		},
		{
			name:     "single number",
			numbers:  []float64{42},
			expected: 0.0, // √0 = 0
		},
		{
			name:     "two numbers",
			numbers:  []float64{1, 5},
			expected: 2.0, // √4 = 2
		},
		{
			name:     "integers with integer std dev",
			numbers:  []float64{1, 2, 3, 4, 5},
			expected: 1.4142135623730951, // √2
		},
		{
			name:     "perfect square variance",
			numbers:  []float64{1, 5, 9},
			expected: 3.265986323710904, // √(32/3) = √10.666...
		},
		{
			name:     "negative numbers",
			numbers:  []float64{-10, -20, -30},
			expected: 8.16496580927726, // √66.666...
		},
		{
			name:     "mixed positive and negative",
			numbers:  []float64{-5, 0, 5},
			expected: 4.08248290463863, // √16.666...
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := standardDeviation(tt.numbers)
			// Use a small delta for float comparison due to precision
			delta := 0.00001
			if result < tt.expected-delta || result > tt.expected+delta {
				t.Errorf("standardDeviation(%v) = %f; expected %f", tt.numbers, result, tt.expected)
			}
		})
	}
}

// TestReadNumbersFromFile tests the file reading functionality
func TestReadNumbersFromFile(t *testing.T) {
	// Table-driven tests for file reading
	tests := []struct {
		name          string      // Description of the test case
		filename      string      // Path to test file
		expected      []float64   // Expected numbers read
		expectError   bool        // Should this test produce an error?
		errorContains string      // Expected error message substring
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

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := readNumbersFromFile(tt.filename)

			// Check error expectation
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

			// Check no error when not expected
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Check the numbers match
			if !floatSlicesEqual(result, tt.expected) {
				t.Errorf("readNumbersFromFile(%s) = %v; expected %v", tt.filename, result, tt.expected)
			}
		})
	}
}

// Helper function to check if two float slices are equal
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

// Helper function to check if a string contains a substring
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
