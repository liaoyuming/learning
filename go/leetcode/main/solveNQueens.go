/**
	n 皇后问题
  	see https://leetcode-cn.com/problems/n-queens/description/
 */

package main

import (
	"fmt"
	"math"
)

func main()  {
	res := solveNQueens(4)
	fmt.Printf("%v", res)
}

func solveNQueens(n int) [][]string {
	pos := make([]int, n)
	var res [][]string
	for i:=0; i<n; i++ {
		pos[i] = -1
	}
	solve(&pos, 0, &res)
	return res
}

func solve(pos *[]int, row int, res *[][]string) {
	n := len(*pos)
	if (row == n) {
		tmp := make([]string, n)
		for r:=0; r<n; r++ {
			for c:=0; c < n; c++ {
				if ((*pos)[r] == c) {
					tmp[r] += "Q"
				} else {
					tmp[r] += "."
				}
			}
		}
		*res = append(*res, tmp)
		return
	}
	for col:=0; col<n; col++ {
		if (isValid(*pos, row, col)) {
			(*pos)[row] = col
			solve(pos, row+1, res)
			(*pos)[row] = -1
		}
	}

}

func isValid(pos []int, row int, col int) bool {
	for i:= 0; i < row; i++ {
		if (pos[i] == col || math.Abs(float64(col - pos[i])) == math.Abs(float64(row - i))) {
			return false
		}
	}
	return true
}