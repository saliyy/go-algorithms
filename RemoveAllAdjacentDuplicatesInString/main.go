package main

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
type Stack struct {
	items []rune
}

func (s *Stack) latest() rune {
	l := len(s.items) - 1
	return s.items[l]
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) pop() rune {
	l := len(s.items) - 1
	removed := s.items[l]
	s.items = s.items[:l]
	return removed
}

func (s *Stack) unshift(el rune) *Stack {
	s.items = append(s.items, el)
	return s
}

func removeDuplicates(s string) string {
	stack := &Stack{items: []rune{}}

	for i := 0; i < len(s); i++ {
		symb := s[i]

		if stack.isEmpty() == false && stack.latest() == rune(symb) {
			stack.pop()
		} else {
			stack.unshift(rune(symb))
		}
	}

	return string(stack.items)
}
