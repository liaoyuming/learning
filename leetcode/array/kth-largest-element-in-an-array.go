package main

func findKthLargest(nums []int, k int) int {
	t := nums[0]
	l := []int{}
	r := []int{}
	for i:=0; i<len(nums); i++ {
		if nums[i] < t {
			l = append(l, nums[i])
		} else if nums[i] > t {
			r = append(r, nums[i])
		}
	}
	if len(r) >= k {
		return findKthLargest(r, k)
	} else if k <= len(nums)-len(l){
		return t
	} else {
		return findKthLargest(l, k-(len(nums)-len(l)))
	}
}
// 堆排
func findKthLargest2(nums []int, k int) int {
	for i:=(len(nums)-1)/2; i>=0; i-- {
		nums = buildHeap(nums, i, len(nums))
	}
	for i:=1; i<k; i++ {
		nums[0], nums[len(nums)-i] = nums[len(nums)-i], nums[0]
		nums = buildHeap(nums, 0, len(nums)-i)
	}
	return nums[0]
}

func buildHeap(nums []int, i int, length int) []int {
	left := 2*i+1
	right := 2*i+2
	large := i
	if left < length && nums[left] > nums[large] {
		large = left
	}
	if right < length && nums[right] > nums[large] {
		large = right
	}
	if large != i {
		nums[i], nums[large] = nums[large], nums[i]
		return buildHeap(nums, large, length)
	}
	return nums
}

// 快排 双指针
func findKthLargest3(nums []int, k int) int {
	t := nums[0]
	left := 0
	right := len(nums)-1
	for left < right {
		for left < right && nums[right] <= t {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] >= t {
			left++
		}
		nums[right] = nums[left]
 	}
 	nums[left] = t
	if left+1 > k {
		return findKthLargest3(nums[:left+1], k)
	} else if k == left+1 {
		return t
	} else {
		return findKthLargest3(nums[left+1:], k-left-1)
	}
}

func main() {
	println(findKthLargest3([]int{3,2,1,5,6,4}, 2))
}