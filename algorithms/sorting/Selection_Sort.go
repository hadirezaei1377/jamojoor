package sorting

import "fmt"

func SelectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {

		minIdx := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main2() {

	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("default array:", arr)
	SelectionSort(arr)
	fmt.Println("sorted array:", arr)
}
