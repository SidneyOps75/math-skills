package Math

import (
	"math"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	testCases := []struct {
		data     []float64
		expected float64
	}{
		{data: []float64{1, 2, 3, 4, 5}, expected: 3.0},
		{data: []float64{10, 20, 30, 40, 50}, expected: 30.0},
		{data: []float64{-1, -2, -3, -4, -5}, expected: -3.0},
	}

	for _, tc := range testCases {
		result := CalculateAverage(tc.data)
		if result != tc.expected {
			t.Errorf("CalculateAverage(%v) = %v; want %v", tc.data, result, tc.expected)
		}
	}
}

func TestCalculateMedian(t *testing.T) {
	testCases := []struct {
		data     []float64
		expected float64
	}{
		{data: []float64{5, 2, 1, 3, 4}, expected: 3.0},
		{data: []float64{6, 2, 1, 3, 4, 5}, expected: 3.5},
		{data: []float64{7, 1, 3, 4, 6}, expected: 4.0},
		{data: []float64{10, 20, 30, 40}, expected: 25.0},
	}

	for _, tc := range testCases {
		result := CalculateMedian(tc.data)
		if result != tc.expected {
			t.Errorf("CalculateMedian(%v) = %v; want %v", tc.data, result, tc.expected)
		}
	}
}

func TestCalculateVariance(t *testing.T) {
	testCases := []struct {
		data     []float64
		mean     float64
		expected float64
	}{
		{data: []float64{1, 2, 3, 4, 5}, mean: 3.0, expected: 2.0},
		{data: []float64{10, 20, 30, 40, 50}, mean: 30.0, expected: 200.0},
		{data: []float64{-1, -2, -3, -4, -5}, mean: -3.0, expected: 2.0},
	}

	for _, tc := range testCases {
		result := CalculateVariance(tc.data, tc.mean)
		if result != tc.expected {
			t.Errorf("CalculateVariance(%v, %v) = %v; want %v", tc.data, tc.mean, result, tc.expected)
		}
	}
}

func TestStandardDeviation(t *testing.T) {
	testCases := []struct {
		data     []float64
		expected float64
	}{
		{data: []float64{1, 2, 3, 4, 5}, expected: math.Sqrt(2.0)},
		{data: []float64{10, 20, 30, 40, 50}, expected: math.Sqrt(200.0)},
		{data: []float64{-1, -2, -3, -4, -5}, expected: math.Sqrt(2.0)},
	}

	for _, tc := range testCases {
		mean := CalculateAverage(tc.data)
		variance := CalculateVariance(tc.data, mean)
		result := math.Sqrt(variance)
		if result != tc.expected {
			t.Errorf("StandardDeviation(%v) = %v; want %v", tc.data, result, tc.expected)
		}
	}
}
