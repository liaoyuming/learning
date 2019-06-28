package main

import "fmt"

/**
 		两数之和 II - 输入有序数组
		https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/description/
 */

func main() {
	fmt.Printf("%v", twoSum([]int{2, 3, 4}, 6))
}

func twoSum(numbers []int, target int) []int {
	len:= len(numbers)
	res := make([]int, 2)
	for i:=0; i<len; i++ {
		num := target - numbers[i]
		start := i+1
		end := len-1
		for start <= end  {
			mid := start + (end - start)/2
			if numbers[mid] < num {
				start = mid + 1
			} else if numbers[mid] > num {
				end = mid - 1
			} else {
				res = []int{i+1, mid+1}
				return res
			}
		}
	}
	return res
}
