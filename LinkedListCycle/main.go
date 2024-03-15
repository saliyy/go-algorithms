package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.com/problems/linked-list-cycle

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	positions := make(map[*ListNode]int)

	startPos := 0

	for head.Next != nil {

		positions[head] = startPos + 1

		_, has := positions[head.Next]

		if has {
			return true
		}

		head = head.Next
	}

	return false
}
