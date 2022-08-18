// https://leetcode.com/problems/shuffle-string/

package main

import "fmt"

func main() {
	fmt.Printf("The origin string is, %s", restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
}

func restoreString(s string, indices []int) string {
	result := make([]byte, len(indices)) // create byte arr

	for index, value := range indices {

		result[value] = s[index]
	}
	fmt.Println(result)
	return string(result)
	// convert byte arr to string
	// [108 101 101 116 99 111 100 101] => leetcode
}

/*
Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
Output: "leetcode"
Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.

*/
