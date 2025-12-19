// Package main provides file I/O operations for the math-skills program.
// This file handles reading numerical data from files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readNumbersFromFile reads numbers from a file and returns them as a slice.
// Each line should contain one number. Invalid lines are skipped silently.
//
// Parameters:
//   - filename: path to the file to read
//
// Returns:
//   - []float64: slice of numbers read from the file
//   - error: error if file cannot be opened or contains no valid numbers
func readNumbersFromFile(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []float64

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue
		}

		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(numbers) == 0 {
		return nil, fmt.Errorf("no valid numbers found in file")
	}

	return numbers, nil
}
