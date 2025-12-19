// Package main implements a statistical analysis tool that calculates
// average, median, variance, and standard deviation from data files.
package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	numbers, err := readNumbersFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	avg := average(numbers)
	med := median(numbers)
	vari := variance(numbers)
	std := standardDeviation(numbers)

	avgRounded := int(math.Round(avg))
	medRounded := int(math.Round(med))
	variRounded := int(math.Round(vari))
	stdRounded := int(math.Round(std))

	fmt.Printf("Average: %d\n", avgRounded)
	fmt.Printf("Median: %d\n", medRounded)
	fmt.Printf("Variance: %d\n", variRounded)
	fmt.Printf("Standard Deviation: %d\n", stdRounded)
}
