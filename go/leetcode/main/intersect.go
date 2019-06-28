package main

import "fmt"

/**
		 两个数组的交集 II
		https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/
 */

func main() {
	fmt.Printf("%v", intersect([]int{1,2,2,1}, []int{2,2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	dist := make(map[int]int, len(nums1))
	var res []int
	for _,v := range nums1 {
		dist[v] += 1
	}
	for _,v := range nums2 {
		if c,ok := dist[v]; ok && c>0{
			res = append(res, v)
			dist[v] -= 1
		}
	}
	return res
}