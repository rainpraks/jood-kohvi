package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input file>")
		return
	}
	inputFileName := os.Args[1]

	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("Error opening input file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []int
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line) // Trim leading/trailing whitespace
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("Error parsing line as integer: %v\n", err)
			}
			data = append(data, num)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}
	average := Average(data)
	fmt.Println("Average:", average)
	median := Median(data)
	fmt.Println("Median:", median)
	variance := Variance(data)
	fmt.Println("Variance:", variance)
	standardDeviation := StandardDeviation(data)
	fmt.Println("Standardeviation:", standardDeviation)
}

func Average(data []int) float64 {
	sum := 0
	count := len(data)
	for _, num := range data {
		sum += num
	}
	average := float64(sum) / float64(count)
	roundedAverage := math.Round(average)
	return roundedAverage
}

func Median(data []int) float64 {
	sort.Ints(data)
	count := len(data)
	if count%2 == 0 {
		median := float64(data[count/2-1]+data[count/2]) / 2.0
		roundedMedian := math.Round(median)
		return roundedMedian
	}
	median := float64(data[count/2])
	roundedMedian := math.Round(median)
	return roundedMedian
}

func Variance(data []int) int {
	count := len(data)
	if count == 0 {
		log.Fatal("Cannot calculate variance for an empty dataset")
	}

	average := Average(data)
	sumOfSquaredDifferences := 0
	for _, num := range data {
		diff := num - int(average)
		sumOfSquaredDifferences += diff * diff
	}
	variance := sumOfSquaredDifferences / count
	return variance
}

func StandardDeviation(data []int) int {
	variance := Variance(data)
	standardDeviation := int(math.Sqrt((float64)(variance)))
	return standardDeviation
}
