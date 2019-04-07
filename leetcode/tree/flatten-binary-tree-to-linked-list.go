package main

import "math/bits"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	right := root.Right
	flatten(right)
	if root.Left != nil {
		root.Right = root.Left
		root.Left = nil
		flatten(root.Right)
		t := root.Right
		for t.Right != nil {
			t = t.Right
		}
		t.Right = right
	}
}