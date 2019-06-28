package main

import (
	"fmt"
)

/**
		平衡二叉树
		https://leetcode-cn.com/problems/balanced-binary-tree/
 */

func main() {
	root := TreeNode{
		1,
		&TreeNode{
			3,
			nil,
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
		nil,
	}
	//head := &ListNode{}
	fmt.Println(isBalanced(&root))
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := true
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *bool) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, res) + 1
	right := dfs(root.Right, res) + 1
	n := 0
	max := 0
	if left > right {
		n = left - right
		max = left
	} else {
		n = right - left
		max = right
	}
	if n > 1 {
		*res = false
	}
	return max
}