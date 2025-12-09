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
