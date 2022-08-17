// https://leetcode.com/problems/water-bottles/

package main

import "fmt"

func main() {
	fmt.Println(numWaterBottles(15, 4))
}

func numWaterBottles(numBottles int, numExchange int) int {

	emptyBottles := numBottles //Initial emptyBottles
	result := numBottles
	var remainEmptyBottles int

	for emptyBottles >= numExchange {
		result += emptyBottles / numExchange
		remainEmptyBottles = emptyBottles % numExchange
		emptyBottles = emptyBottles/numExchange + remainEmptyBottles

	}
	return result

}
