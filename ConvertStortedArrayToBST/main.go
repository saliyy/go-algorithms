package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree
func main() {
	nums := []int{-10, -3, 0, 5, 9}

	fmt.Println(sortedArrayToBST(nums))
}

func sortedArrayToBST(nums []int) *TreeNode {

	var helper func(l, r int) *TreeNode

	helper = func(l, r int) *TreeNode {

		if l <= r {

			mid := (l + r) / 2

			node := &TreeNode{Val: nums[mid]}

			node.Left = helper(l, mid-1)
			node.Right = helper(mid+1, r)

			return node
		}

		return nil
	}

	return helper(0, len(nums)-1)
}
