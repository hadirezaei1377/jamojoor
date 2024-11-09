package lambda

import "fmt"

/* lambda structure in go

func(params) returnType {
    // functions code
}
*/

func main() {
	// sum is a lambda function
	sum := func(a int, b int) int {
		return a + b
	}

	result := sum(5, 7)
	fmt.Println("Result:", result) // Output: Result: 12

	// related to lamdba2.go file :
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// use lambda for even numbers
	evenNumbers := filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println("Even numbers:", evenNumbers) // Output: Even numbers: [2 4 6 8 10]

}
