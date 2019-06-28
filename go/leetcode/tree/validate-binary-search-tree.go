package main

import "github.com/golang/net/html/atom"

type TreeNode struct {
	Val int
	Left *TreeNode
 	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
 	}
 	return help(root.Left, root.Val, -1) && help(root.Right, -1, root.Val)
}

func help(root *TreeNode, max int, min int) bool {
	if root == nil {
		return true
	}
	if (root.Val <= min && min != -1) || (root.Val >= max && max != -1) {
		return false
	}
	return help(root.Left, root.Val, min) && help(root.Right, max, root.Val)
}
