package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(node *TreeNode, path ...string) (out []string) {
	path = append(path, fmt.Sprintf("%v", node.Val))

	if node.Left == nil && node.Right == nil {
		return []string{strings.Join(path, "->")}
	}

	if node.Left != nil {
		out = append(out, binaryTreePaths(node.Left, path...)...)
	}
	if node.Right != nil {
		out = append(out, binaryTreePaths(node.Right, path...)...)
	}

	return
}
