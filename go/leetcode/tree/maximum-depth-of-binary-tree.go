package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	depth := 0
	if left > right {
		depth = left
	} else {
		depth = right
	}
	return depth + 1
}