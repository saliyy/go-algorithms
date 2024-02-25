package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}

//https://leetcode.com/problems/valid-palindrome/submissions/1185780012/

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l <= r {

		if unicode.IsLetter(rune(s[l])) == false && unicode.IsNumber(rune(s[l])) == false {
			l += 1
			continue
		}

		if unicode.IsLetter(rune(s[r])) == false && unicode.IsNumber(rune(s[r])) == false {
			r -= 1
			continue
		}

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}

		l += 1
		r -= 1
	}

	return true
}
