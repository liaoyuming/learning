package main

import (
	"fmt"
	"strconv"
)

/**
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:
   1
 /   \
2     3
 \
  5
输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := []string{}
	help(root, []int{}, &res)
	return res
}

func help(root *TreeNode, tmp []int, res *[]string) {
	t := append(tmp, root.Val)
	if root.Left != nil {
		help(root.Left, t, res)
	}
	if root.Right != nil {
		help(root.Right, t, res)
	}
	if root.Left == nil && root.Right == nil {
		str := ""
		for k,v := range t {
			if k == 0 {
				str += strconv.Itoa(v)
			} else {
				str += "->" + strconv.Itoa(v)
			}
		}
		*res = append(*res, str)
		return
	}
}

func main() {
	a := &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}
	fmt.Println(binaryTreePaths(a))
}