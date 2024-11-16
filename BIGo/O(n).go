package bigo

import "fmt"

func mainn() {

	numbers := []int{10, 20, 30, 40, 50, 60, 70}

	for _, num := range numbers {
		fmt.Println(num)
	}
}
