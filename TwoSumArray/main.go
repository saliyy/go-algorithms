package main

import "fmt"

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted
func main() {
	fmt.Println(twoSum([]int{0, 0, 3, 4}, 0))
}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	sum := 0

	for l <= r {
		sum = numbers[l] + numbers[r]

		if sum == target {
			return []int{l + 1, r + 1}
		}

		if sum < target {
			l = l + 1
		} else {
			r = r - 1
		}
	}

	return []int{}
}
