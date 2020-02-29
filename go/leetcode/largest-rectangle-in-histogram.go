package main

import "fmt"

// 使用单调栈
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := []int{}
	maxArea := 0
	for i:=0; i<len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			length := i
			if len(stack) > 0 {
				length = i-stack[len(stack)-1]-1
			}
			maxArea = max(maxArea, heights[t] * length)
		}
		stack = append(stack, i)
	}
	return maxArea
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
}