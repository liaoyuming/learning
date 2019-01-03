package leetcode

/**
		 打家劫舍
		https://leetcode-cn.com/problems/house-robber-iii/description/
 */

func main() {
	root := w{4, &treenode{1, &treenode{val:2, left:&treenode{val:3}}, nil}, nil}
	//root := treenode{4, &treenode{1, &treenode{val:2}, nil}, nil}
	//root := treenode{}
	println(rob2(&root))
	//println(rob1(&root))
}

 type treenode struct {
 	val int
 	left *treenode
 	right *treenode
 }

 // 递归遍历
func rob(root *treenode) int {
	if root == nil {
		return 0
	}
	num := root.val
	if root.left != nil {
		num += rob(root.left.left) + rob(root.left.right)
	}
	if root.right != nil {
		num += rob(root.right.left) + rob(root.right.right)
	}
	return max(num, rob(root.left) + rob(root.right))
}

// 加了层cache
func rob2(root *treenode) int {
	if root == nil {
		return 0
	}
	cache := map[*treenode]int{}
	dfs(root, &cache)
	return cache[root]
}

func dfs(root *treenode, cache *map[*treenode]int)  {
	if _, ok := (*cache)[root]; ok {
		return
	}
	if root == nil {
		(*cache)[root] = 0
		return
	}
	println(1)
	num := 0
	if root.left != nil {
		dfs(root.left.left, cache)
		dfs(root.left.right, cache)
		num += (*cache)[root.left.left] + (*cache)[root.left.right]
	}
	if root.right != nil {
		dfs(root.right.left, cache)
		dfs(root.right.right, cache)
		num += (*cache)[root.right.left] + (*cache)[root.right.right]
	}
	dfs(root.left, cache)
	dfs(root.right, cache)
	(*cache)[root] = max(root.val + num, (*cache)[root.left] + (*cache)[root.right])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}