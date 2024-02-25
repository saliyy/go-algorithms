package main

import "fmt"

func main() {
	s := []byte("hello")

	reverseString(s)

	fmt.Println(string(s))
}

// https://leetcode.com/problems/reverse-string
func reverseString(s []byte) {
	l, r := 0, len(s)-1

	for l <= r {
		s[l], s[r] = s[r], s[l]
		l += 1
		r -= 1
	}
}
