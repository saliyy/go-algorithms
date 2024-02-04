package main

import "fmt"

/*
https://leetcode.com/problems/student-attendance-record-i/
*/

func main() {
	fmt.Println(checkRecord("PPALLL"))
}

func checkRecord(s string) bool {
	countOfA := 0
	lateInARow := 0

	for _, char := range s {

		if char != 'L' {
			lateInARow = 0
		} else {
			lateInARow += 1
			if lateInARow >= 3 {
				return false
			}
		}

		if char == 'A' {
			countOfA += 1
			if countOfA >= 2 {
				return false
			}
		}
	}

	return true
}
