package main

/**
		存在重复元素 II
		https://leetcode-cn.com/problems/contains-duplicate-ii/description/
 */

func main() {
	println(containsNearbyDuplicate([]int{99,99}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	dist := make(map[int]int, len(nums))
	for i,v := range nums {
		if n, ok := dist[v]; ok {
			if diff(i, n) <= k {
				return true
			}
		}
		dist[v] = i
	}
	return false
}

func diff(a int, b int) int {
	if a > b {
		return a-b
	}
	return b-a
}