package sorting

import "fmt"

func BubbleSort(arr []int) {
	n := len(arr)
	// iterate on all array
	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("default array:", arr)
	BubbleSort(arr)
	fmt.Println("sorted array:", arr)
}
