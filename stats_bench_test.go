// Package main provides statistical calculations for math-skills program.
// This file contains benchmark tests for measuring the performance of
// statistical functions with different dataset sizes.
package main

import (
	"testing"
)

func BenchmarkAverage(b *testing.B) {
	numbers := []float64{189, 113, 121, 114, 145, 110, 99, 134, 156, 178}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = average(numbers)
	}
}

func BenchmarkMedian(b *testing.B) {
	numbers := []float64{189, 113, 121, 114, 145, 110, 99, 134, 156, 178}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = median(numbers)
	}
}

func BenchmarkVariance(b *testing.B) {
	numbers := []float64{189, 113, 121, 114, 145, 110, 99, 134, 156, 178}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = variance(numbers)
	}
}

func BenchmarkStandardDeviation(b *testing.B) {
	numbers := []float64{189, 113, 121, 114, 145, 110, 99, 134, 156, 178}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = standardDeviation(numbers)
	}
}

func BenchmarkReadNumbersFromFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = readNumbersFromFile("testdata/valid_simple.txt")
	}
}

func BenchmarkAverageLargeDataset(b *testing.B) {
	numbers := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = float64(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = average(numbers)
	}
}

func BenchmarkMedianLargeDataset(b *testing.B) {
	numbers := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = float64(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = median(numbers)
	}
}

func BenchmarkVarianceLargeDataset(b *testing.B) {
	numbers := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = float64(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = variance(numbers)
	}
}

func BenchmarkStandardDeviationLargeDataset(b *testing.B) {
	numbers := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = float64(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = standardDeviation(numbers)
	}
}
