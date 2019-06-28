package main

import (
	"fmt"
)

type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generate(1, n, &map[int]map[int][]*TreeNode{})
}

func generate(start int, end int, cache *map[int]map[int][]*TreeNode) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	} else if start == end {
		return []*TreeNode{&TreeNode{Val:start}}
	} else if start == end-1 {
		return []*TreeNode{
			&TreeNode{Val:start, Right:&TreeNode{Val:end}},
			&TreeNode{Val:end, Left:&TreeNode{Val:start}},
		}
	}
	if val, ok := (*cache)[start][end]; ok {
		return val
	}
	res := []*TreeNode{}
	for i:=start; i<=end; i++ {
		left := generate(start, i-1, cache)
		right := generate(i+1, end, cache)
		for _,l := range left {
			for _, r := range right {
				res = append(res, &TreeNode{Val:i, Left:l, Right:r})
			}
		}
	}
	if _, ok := (*cache)[start][end]; !ok {
		(*cache)[start] = map[int][]*TreeNode{}
	}
	(*cache)[start][end] = res
	return res
}

func main() {
	fmt.Println(generateTrees(3))
}
