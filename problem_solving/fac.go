package main

import "fmt"

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main2() {
	fmt.Println("Factorial Calculation:")
	num := 5
	fmt.Printf("%d! = %d\n", num, factorial(num))
}
