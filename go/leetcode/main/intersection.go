package main

import "fmt"

/**
		两个数组的交集
		https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 */

func main() {
	fmt.Printf("%v", intersection([]int{4,9,5}, []int{9,4,9,8,4}))
}

func intersection(nums1 []int, nums2 []int) []int {
	dist1 := make(map[int]int, len(nums1))
	dist2 := make(map[int]int, len(nums2))
	var res []int
	for _,v := range nums1 {
		dist1[v] = 1
	}
	for _,v := range nums2 {
		if _, ok := dist2[v]; !ok {
			dist2[v] = v
			if _,ok := dist1[v]; ok {
				res = append(res, v)
			}
		}
	}
	return res
}