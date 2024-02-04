package main

import (
	"fmt"
)

type Stack struct {
	items []rune
}

func (s *Stack) latest() rune {
	l := len(s.items) - 1
	return s.items[l]
}

func (s *Stack) pop() rune {
	l := len(s.items) - 1
	removed := s.items[l]
	s.items = s.items[:l]
	return removed
}

func isSubsequence(s string, t string) bool {
	stack := Stack{items: []rune(s)}
	runed := []rune(t)
	for i := len(runed) - 1; i >= 0; i-- {
		if len(stack.items) > 0 && stack.latest() == runed[i] {
			stack.pop()
		}
	}

	return len(stack.items) == 0
}

func main() {
	fmt.Println(isSubsequence("abx", "ahbgdc"))
}
