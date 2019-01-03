package leetcode

import "fmt"

func main() {
	var nums = []int{1,1,1,2,2,3,4,4,5}
	fmt.Printf("%v\n", twoSum(nums, 5))
}

func twoSum(nums []int, target int) []int {
	list := make(map[int]int, len(nums))
	for key, val := range nums {
		list[val] = key
	}
	for k, i := range nums{
		if l, ok := list[target-i]; ok && l != k {
			return []int{k, l}
		}
	}
	return nil
}