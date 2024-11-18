package main

import (
	"fmt"
	"sort"
)

func findMax(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func calculateMean(arr []int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

func calculateMedian(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	if n%2 == 0 {
		return float64(arr[n/2-1]+arr[n/2]) / 2
	}
	return float64(arr[n/2])
}

func areArraysEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func main12() {
	array := []int{10, 20, 30, 40, 50}

	fmt.Println("Maximum value:", findMax(array))
	fmt.Println("Mean:", calculateMean(array))
	fmt.Println("Median:", calculateMedian(array))
}
