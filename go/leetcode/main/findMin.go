package main

/**
		寻找旋转排序数组中的最小值
		https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
 */

func main() {
	println(findMin([]int{3,4,5,1,2}))
}

// 使用二分查找
func findMin(nums []int) int {
	len := len(nums)
	start := 0
	end := len-1
	for  end - start > 1 {
		mid := start + (end - start) / 2
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if nums[mid-1] > nums[end] {
			start = mid + 1
		}  else {
			end = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}