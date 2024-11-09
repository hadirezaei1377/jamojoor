package project

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"apple", "banana", "kiwi", "cherry", "blueberry"}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	fmt.Println("Sorted words by length:", words)
}
