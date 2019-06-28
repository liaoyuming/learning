package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 0
	max := 0
	for i:=1; i<len(nums); i++ {
		if nums[i-1] >= nums[i] {
			dp[i] = i
		} else {
			dp[i] = dp[i-1]
		}
		t := i - dp[i] + 1
		if max < t {
			max = t
		}
	}
	return max
}

func main() {
	println(findLengthOfLCIS([]int{1}))
}