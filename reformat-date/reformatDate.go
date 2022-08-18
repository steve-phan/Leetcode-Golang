// https://leetcode.com/problems/reformat-date/

package main

import "fmt"

func main() {
	fmt.Printf("the formatted result is, %s", reformatDate("6th Jun 1933"))
}

func reformatDate(date string) string {
	// create a month Mapping
	monthMapping := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	// Input: date = "6th Jun 1933" => len(date) == 12
	// Input: date = "20th Oct 2052" => len(date) == 13
	if len(date) == 12 {
		date = "0" + date // make sure len(date) is always equal 13
	}

	year := date[9:]                 // slice from 9 - 13
	month := monthMapping[date[5:8]] // slice from 5 - 8, then get values
	day := date[:2]                  // slice from 0 - 2

	return year + "-" + month + "-" + day

}

/*
Example 1:

Input: date = "20th Oct 2052"
Output: "2052-10-20"
Example 2:

Input: date = "6th Jun 1933"
Output: "1933-06-06"
Example 3:

Input: date = "26th May 1960"
Output: "1960-05-26"

*/
