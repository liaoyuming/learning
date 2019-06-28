package main

/**
		打家劫舍 II
		https://leetcode-cn.com/problems/house-robber-ii/
 */

func main() {
	println(rob([]int{1,7,9,2}))
}

func rob(nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	if len == 1 {
		return nums[0]
	} else if len == 2 {
		return max(nums[0], nums[1])
	}
	return max(getMax(nums[1:]), getMax(nums[:len-1]))
}

func getMax(nums []int) int {
	len := len(nums)
	dp := make([]int, len)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i:=2; i<len; i++ {
		dp[i] = max(dp[i-2] + nums[i], dp[i-1])
	}
	return dp[len-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}