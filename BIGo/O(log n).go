package bigo

import "fmt"

// binary search is an example of this time complexity
func binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // middle element

		if arr[mid] == target {
			return true // target is here!
		} else if arr[mid] < target {
			left = mid + 1 // search in right section
		} else {
			right = mid - 1 // search in left section
		}
	}

	return false // element not found
}

func mainlog() {
	// srted array
	numbers := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	target := 7
	if binarySearch(numbers, target) {
		fmt.Println("Found:", target)
	} else {
		fmt.Println("Not Found:", target)
	}
}
