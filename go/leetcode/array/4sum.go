package main

import "fmt"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	sort(&nums, 0, len(nums)-1);
	res := [][]int{}
	a, b := 0, len(nums)-1
	for a < len(nums)-3 {
		b = len(nums)-1
		for b > a + 2{
			help(&res, &nums, a+1, b-1, target-nums[a]-nums[b], a, b)
			for nums[b] == nums[b-1] && b > a + 2 {
				b--
			}
			b--
		}
		for nums[a] == nums[a+1] && a < len(nums)-3{
			a++
		}
		a++
	}
	return res
}

func help(res *[][]int, nums *[]int, start int, end int, target int, a int, b int) {
	if start >= end {
		return
	}
	for start < end {
		t := (*nums)[start] + (*nums)[end]
		if t == target {
			*res = append(*res, []int{(*nums)[a], (*nums)[start], (*nums)[end], (*nums)[b]})
			for (*nums)[start] == (*nums)[start+1] && start < end {
				start++
			}
			for (*nums)[end] == (*nums)[end+1] && start < end {
				end--
			}
			start++
			end--
		} else if t < target {
			for (*nums)[start] == (*nums)[start+1] && start < end {
				start++
			}
			start++
		} else {
			for (*nums)[end] == (*nums)[end+1] && start < end {
				end--
			}
			end--
		}
	}
}


func sort(nums *[]int, start int, end int) {
	if (start >= end) {
		return
	}
	target := (*nums)[start]
	a := start
	b := end
	for a < b {
		for (*nums)[b] >= target && a < b {
			b--
		}
		if a < b {
			(*nums)[a] = (*nums)[b]
		}
		for (*nums)[a] < target && a < b {
			a++
		}
		if a < b {
			(*nums)[b] = (*nums)[a]
		}
	}
	(*nums)[a] = target
	sort(nums, start, a-1)
	sort(nums, a+1, end)
}

func main() {
	fmt.Print(fourSum([]int{5,5,3,5,1,-5,1,-2}, 4))
}