package main

import "math"

/*
https://leetcode.com/problems/middle-of-the-linked-list/submissions/1172180709/
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	valNodeMap := make(map[int]ListNode)

	i := 0
	for head != nil {
		valNodeMap[i] = *head

		i += 1
		head = head.Next
	}

	middleVal := int(math.Ceil(float64(len((valNodeMap)) / 2)))

	startFrom := valNodeMap[middleVal]

	return &startFrom
}

func main() {

}
