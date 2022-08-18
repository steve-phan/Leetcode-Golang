// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/

package main

import "fmt"

func main() {
	fmt.Printf("The result of odd nums is  %d", countOdds(4, 13))

}

func countOdds(low int, high int) int {
	nums := high - low + 1
	result := nums / 2

	if nums%2 == 1 && low%2 == 1 {
		result++
	}

	return result

}

//   4 5 6 7 => 2 odd nums (len == even)
// 3 4 5 6 7 => 2 odd nums (len == odd && low = old )
