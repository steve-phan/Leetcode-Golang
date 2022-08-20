package main

import "fmt"

func main() {
	x := maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"})
	fmt.Printf("max is, %v", x)
}

func maxProduct(words []string) int {
	ints := make([]int, len(words))
	for j, word := range words {
		for i := 0; i < len(word); i++ {
			ints[j] = ints[j] | (1 << uint(word[i]-'a'))
		}
	}

	res := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if ints[i]&ints[j] == 0 {
				res = max(res, len(words[i])*len(words[j]))
			}
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
