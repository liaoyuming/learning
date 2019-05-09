package main

import "github.com/influxdata/influxdb/stress/v2/statement"

type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
 }

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			break
		}
	}
	return root
}