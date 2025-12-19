// Package main provides statistical calculation functions.
// This file contains pure mathematical functions for computing
// average, median, variance, and standard deviation.
package main

import (
	"math"
	"sort"
)

// average calculates the arithmetic mean of a slice of numbers.
//
// Formula: average = (x₁ + x₂ + ... + xₙ) / n
//
// Parameters:
//   - numbers: slice of float64 values
//
// Returns:
//   - float64: the arithmetic mean
func average(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

// median calculates the median (middle value) of a slice of numbers.
// For even counts, it returns the average of the two middle values.
//
// Algorithm:
//  1. Create a copy to avoid modifying the original
//  2. Sort the copy in ascending order
//  3. Return middle element (odd) or average of two middle elements (even)
//
// Parameters:
//   - numbers: slice of float64 values
//
// Returns:
//   - float64: the median value
func median(numbers []float64) float64 {
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)

	n := len(sorted)
	if n%2 == 1 {
		return sorted[n/2]
	}
	mid1 := sorted[n/2-1]
	mid2 := sorted[n/2]
	return (mid1 + mid2) / 2.0
}

// variance calculates the variance (measure of spread) of a slice of numbers.
//
// Formula: variance = Σ(xᵢ - μ)² / n
//
// Algorithm:
//  1. Calculate the average (mean)
//  2. Sum the squared differences from the mean
//  3. Divide by the count
//
// Parameters:
//   - numbers: slice of float64 values
//
// Returns:
//   - float64: the variance
func variance(numbers []float64) float64 {
	avg := average(numbers)

	sumSquaredDiff := 0.0
	for _, num := range numbers {
		diff := num - avg
		sumSquaredDiff += diff * diff
	}

	return sumSquaredDiff / float64(len(numbers))
}

// standardDeviation calculates the standard deviation of a slice of numbers.
// Standard deviation is the square root of variance.
//
// Formula: standardDeviation = √variance
//
// Parameters:
//   - numbers: slice of float64 values
//
// Returns:
//   - float64: the standard deviation
func standardDeviation(numbers []float64) float64 {
	return math.Sqrt(variance(numbers))
}
