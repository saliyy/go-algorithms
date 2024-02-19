package main

import "fmt"

// https://leetcode.com/problems/pascals-triangle
func main() {
	generate(5)
}

func generate(numRows int) [][]int {
	var triangle [][]int

	if numRows == 0 {
		return triangle
	}

	triangle = append(triangle, []int{1})

	for i := 1; i < numRows; i++ {
		prevRow := triangle[i-1]
		var newRow []int
		newRow = append(newRow, 1)

		for j := 1; j < len(prevRow); j++ {
			fmt.Println(prevRow)
			newRow = append(newRow, prevRow[j-1]+prevRow[j])
		}

		newRow = append(newRow, 1)
		triangle = append(triangle, newRow)
	}

	return triangle
}
