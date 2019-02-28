package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	dfs(root, sum, &[]int{}, &result)
	return result
}

func dfs(root *TreeNode, sum int, stash *[]int, result *[][]int)  {
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			*stash = append((*stash)[:len(*stash):len(*stash)], root.Val)
			*result = append(*result, *stash)
			return
		}
	}
	if root.Left != nil {
		*stash = append((*stash)[:len(*stash):len(*stash)], root.Val)
		dfs(root.Left, sum-root.Val, stash, result)
		*stash = (*stash)[:len(*stash)-1]
	}
	if root.Right != nil {
		*stash = append((*stash)[:len(*stash):len(*stash)], root.Val)
		dfs(root.Right, sum-root.Val, stash, result)
		*stash = (*stash)[:len(*stash)-1]
	}
}

func buildTree(nums []int, i int) *TreeNode {
	if nums == nil {
		return nil
	}
	if i > len(nums) {
		return nil
	}
	if nums[i-1] == 0 {
		return nil
	}
	root := &TreeNode{Val:nums[i-1]}
	root.Left = buildTree(nums, 2*i)
	root.Right = buildTree(nums, 2*i+1)
	return root
}

func main() {
	nums := []int{5,4,8,11,0,13,4,7,2,0,0,0,0,5,1}
	root := buildTree(nums, 1)
	//head := &ListNode{}
	fmt.Println(pathSum(root, 22))
}
