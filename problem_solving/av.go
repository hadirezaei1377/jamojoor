package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main5() {
	fmt.Println("Prime Numbers between 1 and 20:")
	for i := 1; i <= 20; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
