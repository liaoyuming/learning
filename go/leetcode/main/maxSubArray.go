package main

/**
		最大子序和
		https://leetcode-cn.com/problems/maximum-subarray/
 */

func main() {
	println(maxSubArray([]int{-1,3,4,-1,-10}))
}

func maxSubArray(nums []int) int {
	len := len(nums)
	sum := 0
	maxSub := nums[0]
	for i:=0; i<len; i++ {
		sum += nums[i]
		maxSub = max(sum, maxSub)
		if sum < 0 {
			sum = 0
		}
	}
	return maxSub
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func sum(nums []int) int {
	count := 0
	for i:=0; i<len(nums); i++ {
		count += nums[i]
	}
	return count
}
