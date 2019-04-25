package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	} else if len(nums) == 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	help(nums, 0, &res)
	return res
}

func help(nums []int, k int, res *[][]int) {
	if k == len(nums) {
		*res = append(*res, append([]int{}, nums...))
		return
	}
	for i:=k; i<len(nums); i++ {
		swap(&nums, k, i)
		help(nums, k+1, res)
		swap(&nums, k, i)
	}
}

func swap(nums *[]int, i int, j int) {
	t := (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = t
}


func permute2(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	} else if len(nums) == 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	for i:=0; i<len(nums); i++ {
		n := append([]int(nil), nums...)
		n = append(n[:i], n[i+1:]...)
		t := permute2(n)
		for j:=0; j<len(t); j++ {
			res = append(res, append(t[j], nums[i]))
		}
	}
	return res
}


func main() {
	res := permute2([]int{1, 2, 3, 4})
	fmt.Print(res, len(res))
}

