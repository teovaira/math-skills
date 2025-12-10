package main

import "sort"

// average calculates the arithmetic mean of a slice of numbers.
// It returns the sum of all numbers divided by the count.
//
// Formula: average = (x₁ + x₂ + ... + xₙ) / n
//
// Parameters:
//   - numbers: slice of float64 values
//
// Returns:
//   - float64: the arithmetic mean
//
// Example:
//
//	average([]float64{1, 2, 3, 4, 5}) returns 3.0
func average(numbers []float64) float64 {
	// Calculate the sum of all numbers
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	// Divide sum by count to get average
	return sum / float64(len(numbers))
}

// median calculates the median (middle value) of a slice of numbers.
// The median is the middle value when numbers are sorted.
// For even counts, it returns the average of the two middle values.
//
// Algorithm:
//  1. Create a copy of the slice to avoid modifying the original
//  2. Sort the copy in ascending order
//  3. If odd count: return the middle element
//  4. If even count: return average of the two middle elements
//
// Parameters:
//   - numbers: slice of float64 values
//
// Returns:
//   - float64: the median value
//
// Example:
//
//	median([]float64{1, 3, 5}) returns 3.0
//	median([]float64{1, 2, 3, 4}) returns 2.5
func median(numbers []float64) float64 {
	// Create a copy to avoid modifying the original slice
	// make() creates a new slice with the same length
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)

	// Sort the copy in ascending order
	sort.Float64s(sorted)

	// Get the count of numbers
	n := len(sorted)

	// Check if count is odd or even
	if n%2 == 1 {
		// Odd count: return the middle element
		// Example: [1, 3, 5] has length 3, middle index is 3/2 = 1
		return sorted[n/2]
	} else {
		// Even count: return average of two middle elements
		// Example: [1, 2, 3, 4] has length 4
		// Middle indices are 4/2-1 = 1 and 4/2 = 2
		// Values are 2 and 3, average is 2.5
		mid1 := sorted[n/2-1]
		mid2 := sorted[n/2]
		return (mid1 + mid2) / 2.0
	}
}

// variance calculates the variance (measure of spread) of a slice of numbers.
// Variance measures how far numbers are spread out from their average.
//
// Formula: variance = Σ(xᵢ - μ)² / n
//
// Algorithm:
//  1. Calculate the average (mean) of the numbers
//  2. For each number, subtract the average and square the result
//  3. Sum all the squared differences
//  4. Divide by the count of numbers
//
// Parameters:
//   - numbers: slice of float64 values
//
// Returns:
//   - float64: the variance
//
// Example:
//
//	variance([]float64{4, 8, 6, 5, 3}) returns 2.96
//	variance([]float64{5, 5, 5, 5}) returns 0.0 (no spread)
func variance(numbers []float64) float64 {
	// Step 1: Calculate the average (reusing our average function)
	avg := average(numbers)

	// Step 2 & 3: Calculate sum of squared differences
	sumSquaredDiff := 0.0
	for _, num := range numbers {
		// Difference from average
		diff := num - avg
		// Square the difference and add to sum
		sumSquaredDiff += diff * diff
	}

	// Step 4: Divide by count to get variance
	return sumSquaredDiff / float64(len(numbers))
}
