package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main8() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Random Numbers:")
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100))
	}
}
