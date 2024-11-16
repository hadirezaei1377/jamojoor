package bigo

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ { // first loop
		for j := 0; j < n-i-1; j++ { // second loop
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func mainn2() {

	numbers := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original Array:", numbers)

	bubbleSort(numbers)

	fmt.Println("Sorted Array:", numbers)
}
