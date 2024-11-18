package main

import "fmt"

func sumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func main7() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", numbers)
	fmt.Printf("Sum: %d\n", sumArray(numbers))
}
