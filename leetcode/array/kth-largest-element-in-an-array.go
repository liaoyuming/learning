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

func main() {
	println(findKthLargest([]int{-1,-1}, 2))
}