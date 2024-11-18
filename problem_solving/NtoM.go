package main

import (
	"fmt"
	"strconv"
)

func convertBase(num string, fromBase, toBase int) (string, error) {
	// N to 10
	decimal, err := strconv.ParseInt(num, fromBase, 64)
	if err != nil {
		return "", err
	}

	// 10 to M
	return strconv.FormatInt(decimal, toBase), nil
}

func main10() {
	num := "1011"
	fromBase := 2
	toBase := 16

	result, err := convertBase(num, fromBase, toBase)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("number %s to %d %s from %d converted\n", num, fromBase, result, toBase)
}
