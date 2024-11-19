package algorithms

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) (int, []int) {

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	count := 0
	selectedCoins := []int{}

	for _, coin := range coins {

		for amount >= coin {
			amount -= coin
			selectedCoins = append(selectedCoins, coin)
			count++
		}
	}

	if amount != 0 {
		return -1, nil
	}
	return count, selectedCoins
}

func main8() {

	coins := []int{1, 5, 10, 25, 50}
	amount := 63

	count, selectedCoins := coinChange(coins, amount)

	if count == -1 {
		fmt.Println("Cannot make the target amount with the given coins.")
	} else {
		fmt.Printf("Minimum number of coins required: %d\n", count)
		fmt.Printf("Selected coins: %v\n", selectedCoins)
	}
}
