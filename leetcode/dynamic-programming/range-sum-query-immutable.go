package main

type NumArray struct {
	dp []int
}


func Constructor(nums []int) NumArray {
	dp := make([]int, len(nums))
	for k, v:= range nums {
		if k == 0 {
			dp[k] = v
		} else {
			dp[k] = dp[k-1] + v
		}
	}
	return NumArray{dp:dp}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.dp[j]
	}
	return this.dp[j] - this.dp[i-1]
}


