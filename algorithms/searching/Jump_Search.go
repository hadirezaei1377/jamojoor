package searching

import (
	"fmt"
	"math"
)

func JumpSearch(arr []int, target int) int {
	n := len(arr)
	step := int(math.Sqrt(float64(n)))
	prev := 0

	for arr[int(math.Min(float64(step), float64(n))-1)] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	for i := prev; i < int(math.Min(float64(step), float64(n))); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main5() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80}
	target := 60

	result := JumpSearch(arr, target)

	if result != -1 {
		fmt.Printf("value %d in index %d founded!\n", target, result)
	} else {
		fmt.Println("value not found!")
	}
}
