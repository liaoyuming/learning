package main

/**
		存在重复元素
		https://leetcode-cn.com/problems/contains-duplicate/description/
 */

func containsDuplicate(nums []int) bool {
	dist := make(map[int]int, len(nums))
	for _,v := range nums {
		if _, ok := dist[v]; ok {
			return true
		}
		dist[v] = 1
	}
	return false
}