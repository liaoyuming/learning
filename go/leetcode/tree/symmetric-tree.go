package main

import "github.com/golang/net/route"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return help(root.Left, root.Right)
}

func help(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return help(root1.Right, root2.Left) && help(root1.Left, root2.Right)
}