package main

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target int, start int, end int) int {
	if end < start {
		return -1
	}
	mid := (end + start)/2
	if nums[mid] > target {
		if nums[mid] > nums[end] && target <= nums[end] {
			return binarySearch(nums, target, mid+1, end)
		}
		return binarySearch(nums, target, start, mid-1)
	} else if target == nums[mid] {
		return mid
	} else {
		if nums[mid] < nums[start] && target >= nums[start] {
			return binarySearch(nums, target, start, mid-1)
		}
		return binarySearch(nums, target, mid+1, end)
	}
}

func main() {
	println(search([]int{4,5,6,7,0,1,2}, 2))
}