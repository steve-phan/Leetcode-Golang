// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/

package main

import "fmt"

func main() {
	fmt.Println(average([]int{3000, 2500, 6000, 1000, 8000}))
}

func average(salary []int) float64 {
	min := salary[0]
	max := salary[0]
	totalSalary := 0

	for _, num := range salary {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		totalSalary += num
	}
	averageNum := float64((totalSalary - (min + max))) / float64(len(salary)-2)

	return averageNum

}
