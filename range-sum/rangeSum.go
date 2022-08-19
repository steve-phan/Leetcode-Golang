// https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/

package main

import "fmt"

func main() {
	sum := rangeSum([]int{1, 2, 3, 4}, 4, 1, 10)
	fmt.Println(sum)
}

func rangeSum(nums []int, n int, left int, right int) int {
	numMap := map[int]int{}
	result := 0

	for i := 0; i < n; i++ {
		numMap[nums[i]]++
		tempNum := nums[i]
		for j := i + 1; j < n; j++ {
			numMap[tempNum+nums[j]]++
			tempNum += nums[j]

		}
	}
	fmt.Println(numMap)
	count := 0
	for k := left; k <= right && count < right+1; k++ {
		fmt.Printf("Current result is  k = %d, numMap[k] = %d,  %d\n", k, numMap[k], result)
		result += (numMap[k] * k)
		count += numMap[k]
	}
	// for key, value := range numMap {
	// 	fmt.Printf("The index and value %d, %v\n", key, value)
	// }
	return result % 1000000007

}

/*
Input: nums = [1,2,3,4], n = 4, left = 1, right = 5
Output: 13
Explanation: All subarray sums are 1, 3, 6, 10, 2, 5, 9, 3, 7, 4. After sorting them in non-decreasing order we have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. The sum of the numbers from index le = 1 to ri = 5 is 1 + 2 + 3 + 3 + 4 = 13.
*/

/*
 [1] [1, 2] [1, 2, 3] [1, 2, 3, 4]  :1 :2 : 3 :4
 [2] [2, 3] [3, 4, 4] 1:2 1:3 1:4
 [3] [3, 4] 2:3 2:4
 [4] 3:4
*/
// Result : 1 2 3 3 4 5 6 7 9 10
// 1 2 6  4  5  6  7  9
//   3 9 13 18 24 33
