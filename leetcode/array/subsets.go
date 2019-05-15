package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) < 1 {
		return [][]int{[]int{}}
	}
	sub := subsets(nums[1:])
	res := sub
	for i:=0; i<len(sub); i++ {
		t := append(sub[i][:len(sub[i]):len(sub[i])], nums[0])
		res = append(res, t)
		fmt.Println(res)
	}
	return res
}

func main() {
	fmt.Print(subsets([]int{5,0,2,3,4}))
}