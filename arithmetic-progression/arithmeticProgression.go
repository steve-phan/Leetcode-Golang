// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/

package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(canMakeArithmeticProgression([]int{1, 3, 5, 7, 9}))

}

func canMakeArithmeticProgression(arr []int) bool {

	// *** WITH sort package****
	sort.Ints(arr)

	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}

	return true

	// ****** WITHOUT sort package *******
	// min := arr[0]
	// max := arr[0]
	// N := len(arr)

	// numsMap := make(map[int]bool, N)
	// // find out min, max and numsMap
	// for _, i := range arr {
	// 	if i < min {
	// 		min = i
	// 	}

	// 	if i > max {
	// 		max = i
	// 	}
	// 	numsMap[i] = true
	// }

	// diffRange := (max - min)

	// if diffRange%(N-1) != 0 {
	// 	return false
	// }

	// diff := diffRange / (N - 1)

	// for min < max {
	// 	min += diff
	// 	if !numsMap[min] {
	// 		return false
	// 	}
	// }

	// return true

}

/*
Example 1:

Input: arr = [3,5,1]
Output: true
Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with differences 2 and -2 respectively, between each consecutive elements.
Example 2:

Input: arr = [1,2,4]
Output: false
Explanation: There is no way to reorder the elements to obtain an arithmetic progression.

*/
