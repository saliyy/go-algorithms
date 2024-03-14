package main

import (
	"fmt"
	"strings"
)

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

func (s *Stack) popright() {
	l := len(s.items) - 1
	s.items = s.items[:l]
}

func (s *Stack) unshift(el rune) *Stack {
	s.items = append(s.items, el)
	return s
}

func remove(s string) string {
	stack := &Stack{items: make([]rune, 0, len(s))}

	for _, letter := range s {
		if string(letter) != "#" {
			stack.unshift(letter)
		} else {
			stack.popright()
		}
	}

	return string(stack.items)

}

func backspaceCompare(s string, t string) bool {
	s = remove(s)
	t = remove(t)

	return strings.Compare(s, t) == 0
}

func main() {
	fmt.Println(backspaceCompare("xywrrmp", "xywrrmu#p"))
}
