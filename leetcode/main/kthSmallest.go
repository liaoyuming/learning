package main

func main() {
	root := TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{6, nil, nil}}
	println(kthSmallest(&root, 3))
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var val int
	kthSmallestIter(root, &k, &val)
	return val
}

func kthSmallestIter(root *TreeNode, k *int, val *int) {
	if root == nil {
		return
	}
	kthSmallestIter(root.Left, k, val)
	*k--
	if *k == 0 {
		*val = root.Val
		return
	}
	kthSmallestIter(root.Right, k, val)
}