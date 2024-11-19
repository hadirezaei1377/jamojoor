package searching

import "fmt"

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2 // fine mid

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main4() {
	arr := []int{10, 20, 30, 40, 50}
	target := 40

	result := BinarySearch(arr, target)

	if result != -1 {
		fmt.Printf("value %d in index %d founded!\n", target, result)
	} else {
		fmt.Println("value not found!")
	}
}
