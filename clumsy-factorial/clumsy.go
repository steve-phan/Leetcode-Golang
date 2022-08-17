// https://leetcode.com/problems/clumsy-factorial/

package main

import "fmt"

func main() {
	fmt.Println(clumsyfactorial(929))
}

func clumsyfactorial(n int) int {

	switch n {
	case 1:
		return 1
	case 2:
		return 2 * 1
	case 3:
		return 3 * 2 / 1
	default:
		return n*(n-1)/(n-2) + t(n-3)
	}
}

func t(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2 - 1
	case 3:
		return 3 - 2*1
	case 4:
		return 4 - 3*2/1
	default:
		return n - (n-1)*(n-2)/(n-3) + t(n-4)
	}
}

/*
Example 1:

Input: n = 4
Output: 7
Explanation: 7 = 4 * 3 / 2 + 1
Example 2:

Input: n = 10
Output: 12
Explanation: 12 = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1

*/
