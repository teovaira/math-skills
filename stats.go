package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

// standardDeviation calculates the standard deviation of a slice of numbers.
// Standard deviation is the square root of variance and measures the typical
// distance of values from the average in the same units as the original data.
//
// Formula: standardDeviation = √variance
//
// Algorithm:
//  1. Calculate the variance (reusing our variance function)
//  2. Return the square root of the variance
//
// Parameters:
//   - numbers: slice of float64 values
//
// Returns:
//   - float64: the standard deviation
//
// Example:
//
//	standardDeviation([]float64{4, 8, 6, 5, 3}) returns ~1.72
//	standardDeviation([]float64{5, 5, 5, 5}) returns 0.0 (no spread)
func standardDeviation(numbers []float64) float64 {
	// Calculate variance using our variance function
	v := variance(numbers)

	// Return the square root of variance
	return math.Sqrt(v)
}

// readNumbersFromFile reads numbers from a file and returns them as a slice.
// Each line in the file should contain one number.
// Invalid lines (non-numeric) are skipped silently.
//
// Algorithm:
//  1. Open the file
//  2. Read line by line
//  3. Try to parse each line as a float64
//  4. Skip invalid lines, collect valid numbers
//  5. Return error if no valid numbers found
//
// Parameters:
//   - filename: path to the file to read
//
// Returns:
//   - []float64: slice of numbers read from the file
//   - error: error if file cannot be opened or contains no valid numbers
//
// Example:
//
//	numbers, err := readNumbersFromFile("data.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
func readNumbersFromFile(filename string) ([]float64, error) {
	// Step 1: Open the file
	file, err := os.Open(filename)
	if err != nil {
		// Return error if file cannot be opened
		// This covers "file does not exist" case
		return nil, err
	}
	// Defer closing the file to ensure it's closed when function exits
	defer file.Close()

	// Step 2: Create a scanner to read line by line
	scanner := bufio.NewScanner(file)

	// Slice to store valid numbers
	var numbers []float64

	// Step 3: Read each line
	for scanner.Scan() {
		line := scanner.Text()

		// Trim whitespace from line
		line = strings.TrimSpace(line)

		// Skip empty lines
		if line == "" {
			continue
		}

		// Step 4: Try to parse as float64
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			// Skip invalid lines silently
			// This allows files with mixed valid/invalid data
			continue
		}

		// Add valid number to slice
		numbers = append(numbers, num)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Step 5: Check if we got at least one valid number
	if len(numbers) == 0 {
		return nil, fmt.Errorf("no valid numbers found in file")
	}

	return numbers, nil
}
