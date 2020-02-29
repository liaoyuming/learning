package main

import "os"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	help(root, 0, &res)
	return res
}

func help(root *TreeNode, deep int, res *[]int)  {
	if root == nil {
		return
	}
	deep++
	if deep > len(*res) {
		*res = append(*res, root.Val)
	}
	help(root.Right, deep, res)
	help(root.Left, deep, res)
}