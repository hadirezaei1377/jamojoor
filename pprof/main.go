package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go startCPULoad()
	go startMemoryLoad()

	fmt.Println("Server is running at :6060")
	http.ListenAndServe(":6060", nil)
}

func startCPULoad() {
	for {
		calculatePrimes(5000)
		time.Sleep(500 * time.Millisecond)
	}
}

func startMemoryLoad() {
	for {
		allocateMemory()
		time.Sleep(1 * time.Second)
	}
}

func calculatePrimes(limit int) {
	primes := []int{}
	for i := 2; i < limit; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
}

func allocateMemory() {
	_ = make([]byte, 10*1024*1024)
}
