package main

import "fmt"

/**
		杨辉三角 II
		https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 */

func main() {
	fmt.Printf("%v", getRow(3))
}

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	a := 0
	b := 0
	for i:=0; i<rowIndex+1; i++ {
		if i == 0 {
			res[0] = 1
			continue
		}
		for j:=1; j<i+1; j++ {
			if j == 1 {
				a = res[0]
			} else {
				a = b
			}
			b = res[j]
			if j == i {
				res[j] = a
			} else {
				res[j] = a + b
			}
		}
	}
	return res
}