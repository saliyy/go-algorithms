package main

import (
	"fmt"
)

//https://leetcode.com/problems/binary-search/

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}

	fmt.Println(search(nums, 9))
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		middle := (left + right) / 2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			right = right - 1
		} else {
			left = left + 1
		}
	}

	return -1
}
