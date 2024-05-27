package Math

import (
	"math"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	expected := 3.0
	result := CalculateAverage(data)
	if result != expected {
		t.Errorf("calculateAverage(%v) = %v; want %v", data, result, expected)
	}
}

func TestCalculateMedian(t *testing.T) {
	dataOdd := []float64{5, 2, 1, 3, 4}
	expectedOdd := 3.0
	resultOdd := CalculateMedian(dataOdd)
	if resultOdd != expectedOdd {
		t.Errorf("calculateMedian(%v) = %v; want %v", dataOdd, resultOdd, expectedOdd)
	}

	dataEven := []float64{6, 2, 1, 3, 4, 5}
	expectedEven := 3.5
	resultEven := CalculateMedian(dataEven)
	if resultEven != expectedEven {
		t.Errorf("calculateMedian(%v) = %v; want %v", dataEven, resultEven, expectedEven)
	}
}

func TestCalculateVariance(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	mean := 3.0
	expected := 2.0
	result := CalculateVariance(data, mean)
	if result != expected {
		t.Errorf("calculateVariance(%v, %v) = %v; want %v", data, mean, result, expected)
	}
}

func TestStandardDeviation(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	mean := CalculateAverage(data)
	variance := CalculateVariance(data, mean)
	expected := math.Sqrt(variance)
	result := math.Sqrt(CalculateVariance(data, mean))
	if result != expected {
		t.Errorf("StandardDeviation(%v) = %v; want %v", data, result, expected)
	}
}
