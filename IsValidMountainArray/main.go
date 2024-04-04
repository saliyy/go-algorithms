package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{3, 5, 5}))
}

func validMountainArray(arr []int) bool {
	up, down := false, false
	for i := 1; i < len(arr); i++ {
		curr, prev := arr[i], arr[i-1]
		if !down && curr > prev {
			up = true
			continue
		}
		if up && curr < prev {
			down = true
			continue
		}
		return false
	}
	if up && down {
		return true
	}
	return false
}
