package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) < 1 {
		return []int{}
	} else if len(matrix) < 2 {
		return matrix[0]
	}
	left := 0
	right := len(matrix[0])-1
	top := 0
	bottom := len(matrix)-1
	res := []int{}
	for top <= bottom && left <= right {
		for i:=left; i<=right; i++ {
			res = append(res, matrix[top][i])
		}
		top ++
		for i:=top; i<=bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if top <= bottom {
			for i:=right; i>=left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i:=bottom; i>=top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}