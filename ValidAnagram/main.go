package main

import "fmt"

func main() {
	fmt.Println(isAnagram("a", "ab"))
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sLetters := countLetters(s)
	tLetters := countLetters(t)

	for k, _ := range sLetters {

		if sLetters[k] != tLetters[k] {
			return false
		}
	}

	return true
}

func countLetters(s string) map[rune]int {
	sLettersCount := make(map[rune]int)

	for i := 0; i < len(s); i++ {
		_, has := sLettersCount[rune(s[i])]

		if !has {
			sLettersCount[rune(s[i])] = 1
		} else {
			sLettersCount[rune(s[i])] = sLettersCount[rune(s[i])] + 1
		}
	}

	return sLettersCount
}
