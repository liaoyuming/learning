package leetcode

import "fmt"

/**
		二叉搜索树结点最小距离

给定一个二叉搜索树的根结点 root, 返回树中任意两节点的差的最小值。

示例：

输入: root = [4,2,6,1,3,null,null]
输出: 1
解释:
注意，root是树结点对象(TreeNode object)，而不是数组。

给定的树 [4,2,6,1,3,null,null] 可表示为下图:

          4
        /   \
      2      6
     / \
    1   3

最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
注意：

二叉树的大小范围在 2 到 100。
二叉树总是有效的，每个节点的值都是整数，且不重复。

https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/description/
 */

func main() {
	//arr := []int{4,2,6,1,3, nil, nil}
	root := TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{6, nil, nil}}
	println(minDiffInBST(&root))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	var arr []int
	traverseBST(root, &arr)
	fmt.Printf("%v\n", arr)
	minDiff := 0
	for i:=1; i<len(arr); i++ {
		if i == 1 {
			minDiff = diff(arr[i], arr[i-1])
		}
		minDiff = min(diff(arr[i], arr[i-1]), minDiff)
	}
	return minDiff
}
// 中序遍历
func traverseBST(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	traverseBST(root.Left, arr)
	*arr = append(*arr, root.Val)
	traverseBST(root.Right, arr)
}

func diff(a int, b int) int {
	if a > b {
		return a-b
	}
	return b-a
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}