// https://leetcode.com/problems/number-of-good-pairs/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 1, 1, 3}
	arr1 := []int{1, 1, 1, 1}
	fmt.Printf("The result is: %d\n", numOfGF(arr))
	fmt.Printf("The result is: %d", numOfGF(arr1))

	fmt.Printf("\nThe result is: %d\n", mapNumOfGF(arr))
	fmt.Printf("The result is: %d", mapNumOfGF(arr1))

}

// O(N^2) <= Time complexity
// O(1)   <= Space complexity
func numOfGF(arr []int) int {
	var result int

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				result++
			}
		}
	}

	return result

}

// O(N) <= Time complexity
// O(N) <= Space complexity
func mapNumOfGF(arr []int) int {
	result := 0
	numMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		count, ok := numMap[arr[i]]
		if ok {
			numMap[arr[i]]++
			result += count
		} else {
			numMap[arr[i]] = 1
		}
	}
	return result
}

/*
Example 1:

Input: nums = [1,2,3,1,1,3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
Example 2:

Input: nums = [1,1,1,1]
Output: 6
Explanation: Each pair in the array are good.
Example 3:

Input: nums = [1,2,3]
Output: 0

*/
