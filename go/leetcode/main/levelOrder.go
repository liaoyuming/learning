package main

import "fmt"

/**
		二叉树的层次遍历
		https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 */

func main() {
	root := TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{6, nil, nil}}
	//head := &ListNode{}
	fmt.Printf("%v", levelOrder(&root))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
 
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	dfs(root, &res, 0)
	return res
}

func dfs(root *TreeNode, res *[][]int, deep int) {
	if root == nil {
		return
	}
	dfs(root.Left, res, deep+1)
	if len(*res) <= deep {
		for i:=len(*res); i<=deep; i++ {
			*res = append(*res, []int{})
		}
	}
	(*res)[deep] = append((*res)[deep], root.Val)
	dfs(root.Right, res, deep+1)
}