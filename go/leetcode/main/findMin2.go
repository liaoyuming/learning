package main

/**
		寻找旋转排序数组中的最小值 II
		https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/description/
 */

func main() {
	println(findMin([]int{3,1}))
}

// 使用二分查找
func findMin(nums []int) int {
	len := len(nums)
	start := 0
	end := len-1
	if nums[start] < nums[end] {
		return nums[start]
	}
	for  end >= start {
		mid := start + (end - start) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		}  else if nums[mid] < nums[end] {
			end = mid
		} else {
			end--
		}
	}
	return nums[start]
}