package main

import "github.com/go-kit/kit/metrics/multi"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val:nums[0]}
	}
	mid := len(nums) / 2
	root := &TreeNode{Val:nums[mid]}
	if mid > 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}
	if mid < len(nums)-1 {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}
