// https://leetcode.com/problems/path-crossing/
package main

import "fmt"

func main() {
	fmt.Println(isPathCrossing("WES"))

}

func isPathCrossing(path string) bool {

	// WOW golang can take array as a key of map WOW
	locationsMap := make(map[[2]int]byte)
	currentLocaltion := [2]int{0, 0}
	locationsMap[currentLocaltion] = 0
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'N':
			currentLocaltion[0] += 1
		case 'E':
			currentLocaltion[1] += 1
		case 'S':
			currentLocaltion[0] -= 1
		case 'W':
			currentLocaltion[1] -= 1
		}
		_, ok := locationsMap[currentLocaltion]
		// Check current location is existed or not
		if ok {
			return true
		}
		// assign to map if it's not exist
		locationsMap[currentLocaltion] = 0
	}
	return false
}
