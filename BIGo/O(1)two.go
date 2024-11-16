package bigo

import "fmt"

func maintwo() {

	ageMap := map[string]int{
		"Ali":  30,
		"Sara": 25,
		"Reza": 40,
	}

	fmt.Println("Sara's age:", ageMap["Sara"]) // output: 25
}
