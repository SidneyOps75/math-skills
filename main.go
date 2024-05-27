package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	Math "math-skills/skills"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <datafile>")
		return
	}

	// Open the file
	openFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer openFile.Close()

	// Read lines from the file
	var lines []string
	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// Convert lines to float64 slice
	data := make([]float64, len(lines))
	for i, line := range lines {
		num, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		data[i] = num
	}

	// Calculate statistics
	average := Math.CalculateAverage(data)
	median := Math.CalculateMedian(data)
	variance := Math.CalculateVariance(data, average)
	stdDev := math.Sqrt(variance)

	// Print the results
	fmt.Printf("Average: %d\n", int(math.Round(average)))
	fmt.Printf("Median: %d\n", int(math.Round(median)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}
