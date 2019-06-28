package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val:rootVal}
	if len(inorder) == 1 {
		return root
	}
	rootIndex := 0
	for rootIndex<len(inorder) {
		if rootVal == inorder[rootIndex] {
			break
		}
		rootIndex++
	}
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}