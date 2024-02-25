package main

import "fmt"

func main() {

	arr := []int{3, 4, 5, 1}
	fmt.Println(peakIndexInMountainArray(arr))
}

// https://leetcode.com/problems/peak-index-in-a-mountain-array/submissions/1185745838/
func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := (left + right) / 2

		if arr[middle] > arr[middle-1] && arr[middle] > arr[middle+1] {
			return middle
		}

		if arr[middle] < arr[middle+1] {
			left += 1
		} else {
			if arr[middle] < arr[middle-1] {
				right -= 1
			}
		}
	}

	return -1
}
