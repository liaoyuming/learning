package main

 type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
 }

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	help(root, &max)
	return max
}

func help(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := help(root.Left, max)
	right := help(root.Right, max)
	if left+right > *max {
		*max = left+right
	}
	if left > right {
		return left+1
	} else {
		return right+1
	}
}