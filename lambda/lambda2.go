package lambda

func filter(numbers []int, condition func(int) bool) []int {
	var result []int
	for _, number := range numbers {
		if condition(number) {
			result = append(result, number)
		}
	}
	return result
}
