package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func mainn() {
	fmt.Println("Fibonacci Series:")
	for i := 0; i < 10; i++ {
		fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
	}
}
