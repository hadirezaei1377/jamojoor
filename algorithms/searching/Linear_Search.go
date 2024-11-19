package searching

import "fmt"

func LinearSearch(arr []int, target int) int {

	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func main6() {
	arr := []int{10, 20, 30, 40, 50}
	target := 30

	result := LinearSearch(arr, target)

	if result != -1 {
		fmt.Printf("value %d in index %d founded!\n", target, result)
	} else {
		fmt.Println("value not found!")
	}
}
