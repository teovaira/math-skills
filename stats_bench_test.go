package main

import (
	"testing"
)

// Benchmark for average function
func BenchmarkAverage(b *testing.B) {
	numbers := []float64{189, 113, 121, 114, 145, 110, 99, 134, 156, 178}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = average(numbers)
	}
}

// Benchmark for median function
func BenchmarkMedian(b *testing.B) {
	numbers := []float64{189, 113, 121, 114, 145, 110, 99, 134, 156, 178}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = median(numbers)
	}
}

// Benchmark for variance function
func BenchmarkVariance(b *testing.B) {
	numbers := []float64{189, 113, 121, 114, 145, 110, 99, 134, 156, 178}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = variance(numbers)
	}
}

// Benchmark for standard deviation function
func BenchmarkStandardDeviation(b *testing.B) {
	numbers := []float64{189, 113, 121, 114, 145, 110, 99, 134, 156, 178}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = standardDeviation(numbers)
	}
}

// Benchmark for file reading with different sizes
func BenchmarkReadNumbersFromFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = readNumbersFromFile("testdata/valid_simple.txt")
	}
}

// Benchmark with larger dataset
func BenchmarkAverageLargeDataset(b *testing.B) {
	// Generate 1000 numbers
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
	// Generate 1000 numbers
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
	// Generate 1000 numbers
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
	// Generate 1000 numbers
	numbers := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = float64(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = standardDeviation(numbers)
	}
}
