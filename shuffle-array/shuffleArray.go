// https://leetcode.com/problems/shuffle-the-array/

package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}

func shuffle(nums []int, n int) []int {

	newNums := make([]int, len(nums))

	for i := 0; i < n; i++ {

		//1tim: fill a couple Xi Yi
		// i = 0 => filled 0 , 1 positions
		// i = 1 => filled 2 , 3 positions
		// i = 2 => filled 4 , 5 positions
		// and so on ...

		newNums[i*2] = nums[i]
		newNums[i*2+1] = nums[n+i]
	}
	return newNums

}

/*
xample 1:

Input: nums = [2, 5, 1, 3, 4, 7], n = 3
               x1 x2 x3 y1 y2 y3
Output: [2,3,5,4,1,7]
Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the answer is [2,3,5,4,1,7]  => [x0, y0, x1, y1, x2, y2, x3, y3]
i  = 0 => x0 = nums[0]
          y0 = nums[n]
i  = 1 => x1 = nums[1]
          y1 = nums[n + 1]

Example 2:

Input: nums = [1,2,3,4,4,3,2,1], n = 4
Output: [1,4,2,3,3,2,4,1]
Example 3:

Input: nums = [1,1,2,2], n = 2
Output: [1,2,1,2]

*/
