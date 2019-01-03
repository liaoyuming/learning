package main

/**
		存在重复元素 III
		https://leetcode-cn.com/problems/contains-duplicate-iii/description/
 */

func main() {
	println(containsNearbyAlmostDuplicate([]int{1,0,1,1},1,2))
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	len := len(nums)
	for i,_ := range nums {
		for j:=i+1; j<=i+k && j<len; j++{
			if diff(nums[i], nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func diff(a int, b int) int {
	if a > b {
		return a-b
	}
	return b-a
}