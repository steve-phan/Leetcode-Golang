package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	// Create a DP array to store the minimum number of coins for each amount
	dp := make([]int, amount+1)
	
	// Initialize the DP array with a value larger than the maximum possible number of coins
	for i := range dp {
		dp[i] = amount + 1
	}
	
	// Base case: 0 coins are needed to make amount 0
	dp[0] = 0
	
	// Fill the DP array
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
		}
	}
	
	// Check if the amount can be made up with the given coins
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11

	result := coinChange(coins, amount)
	fmt.Println("Minimum coins needed:", result)
}
