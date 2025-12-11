package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	// Step 1: Check command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <filename>")
		os.Exit(1)
	}

	// Step 2: Get filename from arguments
	filename := os.Args[1]

	// Step 3: Read numbers from file
	numbers, err := readNumbersFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Step 4: Calculate all statistics
	avg := average(numbers)
	med := median(numbers)
	vari := variance(numbers)
	std := standardDeviation(numbers)

	// Step 5: Round to nearest integer
	avgRounded := int(math.Round(avg))
	medRounded := int(math.Round(med))
	variRounded := int(math.Round(vari))
	stdRounded := int(math.Round(std))

	// Step 6: Print results in the required format
	fmt.Printf("Average: %d\n", avgRounded)
	fmt.Printf("Median: %d\n", medRounded)
	fmt.Printf("Variance: %d\n", variRounded)
	fmt.Printf("Standard Deviation: %d\n", stdRounded)
}
