package main

import "fmt"

/**
		二叉树的锯齿形层次遍历
		https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				4,
				nil,
				nil,
			},
			nil,
		},
		&TreeNode{
			3,
			nil,
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
	}
	//head := &ListNode{}
	fmt.Printf("%v", zigzagLevelOrder(&root))
}


func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	traverse(root, &res, 0)
	return res
}

func  traverse (root *TreeNode, res *[][]int, deep int) {
	if root == nil {
		return
	}
	traverse(root.Left, res, deep+1)
	traverse(root.Right, res, deep+1)
	if len(*res) <= deep {
		for i:=len(*res); i<=deep; i++ {
			*res = append(*res, []int{})
		}
	}
	if deep%2 == 1 {
		(*res)[deep] = append([]int{root.Val}, (*res)[deep]...)
	} else {
		(*res)[deep] = append((*res)[deep], root.Val)
	}
}