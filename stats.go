package main

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
