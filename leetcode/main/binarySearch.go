package main

func main() {
	println(search([]int{5}, 5))
}

func search(nums []int, target int) int {
	//return search2(nums, 0, len(nums)-1, target)
	len := len(nums)
	start := 0
	end := len-1
	for start <= end  {
		mid := start + (end - start) / 2
		if  nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func search2(nums []int, start int, end int, target int) int {
	if start > end {
		return -1
	}
	mid := start + (end - start) / 2
	if nums[mid] < target {
		return search2(nums, mid+1, end, target)
	} else if nums[mid] > target {
		return search2(nums, end, mid-1, target)
	} else {
		return mid
	}
}