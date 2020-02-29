package main

import "fmt"

func nextPermutation(nums []int)  {
	if len(nums) < 2 {
		return
	}
	i := len(nums)-2
	for i>=0 && nums[i] >= nums[i+1] {
		i--
	}
	if i == -1 {
		help(0, len(nums)-1, nums)
		return
	}
	j := i
	for j<len(nums)-1 && nums[i] < nums[j+1] {
		j++
	}
	nums[i], nums[j] = nums[j], nums[i]
	help(i+1, len(nums)-1, nums)
}

func help(start int, end int, nums []int) {
	for start < end {
	nums[start], nums[end] = nums[end], nums[start]
	start++
	end--
	}
}

func main() {
	nums := []int{5, 1, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
