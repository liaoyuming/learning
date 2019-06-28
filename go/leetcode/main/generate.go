package main

import "fmt"

/**
		杨辉三角
		https://leetcode-cn.com/problems/pascals-triangle/description/
 */

func main() {
	fmt.Printf("%v", generate(200))
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i:=0; i<numRows; i++ {
		res[i] = make([]int, i+1)
		if i == 0 {
			res[0][0] = 1
			continue
		}
		for j:=0; j<i+1; j++ {
			if j == 0 {
				res[i][j] = res[i-1][0]
			} else if j == i {
				res[i][j] = res[i-1][i-1]
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}