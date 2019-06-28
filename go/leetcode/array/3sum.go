package main

import "fmt"

func threeSum(nums []int) [][]int {
	if len(nums)  < 3 {
		return nil
	}
	sort(&nums, 0, len(nums)-1)
	res := [][]int{}
	for i:=0; i<len(nums)-2 && nums[i] <= 0; i++ {
		target := -nums[i]
		a := i+1
		b := len(nums)-1
		for a < b {
			if nums[a] + nums[b] > target {
				b--
			} else if nums[a] + nums[b] < target {
				a++
			} else {
				res = append(res, []int{nums[i], nums[a], nums[b]})
				for a < b && nums[a+1] == nums[a] {
					a++
				}
				a++
			}
		}
		for i<len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func sort(arr *[]int, start int, end int) {
	if start >= end {
		return
	}
	t := (*arr)[start]
	a := start
	b := end
	for a < b {
		for a < b && (*arr)[b] >= t {
			b --
		}
		if a < b {
			(*arr)[a] = (*arr)[b]
		}
		for a < b && (*arr)[a] < t {
			a ++
		}
		if a < b {
			(*arr)[b] = (*arr)[a]
		}
	}
	(*arr)[a] = t
	sort(arr, start, a-1)
	sort(arr, a+1, end)
}

func main() {
	fmt.Print(threeSum([]int{0, 0, 0}))
}