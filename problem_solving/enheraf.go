package main

import (
	"fmt"
	"math"
)

func calculateStandardDeviation(arr []float64) float64 {
	n := float64(len(arr))
	mean := 0.0
	for _, value := range arr {
		mean += value
	}
	mean /= n

	variance := 0.0
	for _, value := range arr {
		variance += math.Pow(value-mean, 2)
	}
	variance /= n

	return math.Sqrt(variance)
}

func main11() {
	array := []float64{10, 20, 30, 40, 50}
	fmt.Println("Standard Deviation:", calculateStandardDeviation(array))
}
