package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 0, 0}}))
}

func flipAndInvertImage(image [][]int) [][]int {
	for _, sector := range image {
		slices.Reverse(sector)
		for i := 0; i < len(sector); i++ {
			sector[i] = Btoi(sector[i])
		}
	}

	return image
}

func Btoi(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}
