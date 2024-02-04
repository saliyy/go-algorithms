package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Res struct {
	res []int
}

func main() {

}

func (r *Res) traversal(node *TreeNode) {
	if node != nil {
		r.traversal(node.Left)
		r.traversal(node.Right)
		r.res = append(r.res, node.Val)
	}
}

/**
 * Definition for a binary tree node.
 *
 */
func postorderTraversal(root *TreeNode) []int {
	res := Res{}
	res.traversal(root)
	return res.res
}
