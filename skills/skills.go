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
	return sum / float64(len(data))
}

func CalculateMedian(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	}
	return data[n/2]
}

func CalculateVariance(data []float64, mean float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += math.Pow(value-mean, 2)
	}
	return sum / float64(len(data))
}
