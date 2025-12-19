package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
