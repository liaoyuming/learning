package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	maxDP := make([]int, len(nums))
	minDP := make([]int, len(nums))
	maxDP[0] = nums[0]
	minDP[0] = nums[0]
	result := nums[0]
	for i:=1; i<len(nums); i++{
		maxDP[i] = max(nums[i], maxDP[i-1]*nums[i], minDP[i-1]*nums[i])
		minDP[i] = min(nums[i], maxDP[i-1]*nums[i], minDP[i-1]*nums[i])
		if result < maxDP[i] {
			result = maxDP[i]
		}
	}
	return result
}

func max(a int, b int, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func min(a int, b int, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func main() {
	fmt.Println(maxProduct([]int{-1,-1,-2,-2}))
}