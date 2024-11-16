package bigo

import "fmt"

func mainntwo() {

	numbers := []int{10, 20, 30, 40, 50}

	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Println("Sum of elements:", sum)
}
