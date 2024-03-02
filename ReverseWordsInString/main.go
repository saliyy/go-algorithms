package main

import (
	"bytes"
	"fmt"
	"strings"
)

// https://leetcode.com/problems/reverse-words-in-a-string-iii/
func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

// "Let's take LeetCode contest"
func reverseWords(s string) string {
	words := (strings.Split(s, " "))

	var res bytes.Buffer

	for wI, w := range words {
		for i := len(w) - 1; i >= 0; i-- {
			res.WriteRune(rune(w[i]))
		}

		if wI != len(words)-1 {
			res.WriteString(" ")
		}
	}

	return res.String()
}
