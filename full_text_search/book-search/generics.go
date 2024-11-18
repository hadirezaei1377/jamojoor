package booksearch

import "fmt"

// generic function
func reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func mainn() {
	// generic for int
	ints := []int{1, 2, 3, 4}
	fmt.Println("Reversed ints:", reverse(ints))

	// generic for string
	strings := []string{"A", "B", "C", "D"}
	fmt.Println("Reversed strings:", reverse(strings))
}
