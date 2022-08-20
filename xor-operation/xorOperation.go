// https://leetcode.com/problems/xor-operation-in-an-array/

package main

import "fmt"

func main() {
	fmt.Println(xorOperation(19, 8))
}

func xorOperation(n int, start int) int {
	result := start
	for i := 1; i < n; i++ {
		// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Bitwise_XOR
		result ^= start + 2*i
	}

	return result
}
