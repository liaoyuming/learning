package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	root := &TreeNode{Val:postorder[len(postorder)-1]}
	i := 0
	for i<len(inorder) {
		if root.Val == inorder[i] {
			break
		}
		i++
	}
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	root.Left = buildTree(inorder[:i], postorder[:i])
	return root
}