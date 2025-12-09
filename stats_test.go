package main

import "testing"

// TestAverage tests the average calculation function
func TestAverage(t *testing.T) {
	// Test case 1: Simple positive numbers
	numbers := []float64{1, 2, 3, 4, 5}
	expected := 3.0

	result := average(numbers)

	if result != expected {
		t.Errorf("average(%v) = %f; expected %f", numbers, result, expected)
	}
}
