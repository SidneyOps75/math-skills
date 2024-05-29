package Math

import (
	"math"
	"sort"
)

func CalculateAverage(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	average := float64(sum) / float64(len(data))
	return average
}

func CalculateMedian(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	} else {
		return data[n/2]
	}
}

func CalculateVariance(data []float64, average float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += math.Pow(value-average, 2)
	}
	return sum / float64(len(data))
}
