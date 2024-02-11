package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var answer bool

func isSymmetric(root *TreeNode) bool {
	answer = true
	if root != nil {
		recurseSymmetric(root.Left, root.Right)
	}
	return answer
}

func recurseSymmetric(root1, root2 *TreeNode) {
	if root1 == nil && root2 == nil || answer == false {
		return
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		answer = false
		return
	}
	recurseSymmetric(root1.Left, root2.Right)
	recurseSymmetric(root1.Right, root2.Left)
}
