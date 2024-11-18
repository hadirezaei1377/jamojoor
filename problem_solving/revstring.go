package main

import "fmt"

func reverseString(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main6() {
	str := "Hello, Go!"
	fmt.Printf("Original String: %s\n", str)
	fmt.Printf("Reversed String: %s\n", reverseString(str))
}
