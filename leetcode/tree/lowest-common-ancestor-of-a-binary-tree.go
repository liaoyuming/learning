package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l1 := []*TreeNode{}
	l2 := []*TreeNode{}
	l1 = searchOne(root, p, l1)
	l2 = searchOne(root, q, l2)
	i :=1
	for ; i<len(l1) && i<len(l2); i++{
		if l1[i] != l2[i] {
			 return l1[i-1]
		}
	}
	return l1[i-1]
}

func searchOne(root *TreeNode, target *TreeNode, l []*TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	l = append(l, root)
	if root == target {
		return l
	}
	t := searchOne(root.Left, target, l)
	if len(t) > 0 {
		return t
	}
	return searchOne(root.Right, target, l)
}

// 方法2
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}


func main() {
	a := &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}
	b := &TreeNode{3, &TreeNode{5, nil, nil}, a}
	root := TreeNode{1, nil, b}
	r := lowestCommonAncestor2(&root, a, b);
	println(r.Val)
}