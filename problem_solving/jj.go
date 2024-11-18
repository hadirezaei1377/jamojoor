package main

import "fmt"

func main3() {
	a, b := 10, 20
	fmt.Printf("Before Swap: a = %d, b = %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("After Swap: a = %d, b = %d\n", a, b)
}
