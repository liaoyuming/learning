package main

func generateMatrix(n int) [][]int {
	if n < 1 {
		return [][]int{}
	} else if n == 1 {
		return [][]int{[]int{1}}
	}
	res := make([][]int, n)
	for i:=0; i<n; i++ {
		res[i] = make([]int, n)
	}
	left := 0
	right := n-1
	top := 0
	bottom := n-1
	val := 1
	for top <= bottom && left <= right {
		for i:=left; i<=right; i++ {
			res[top][i] = val
			val++
		}
		top++
		for i:=top; i<=bottom; i++ {
			res[i][right] = val
			val++
		}
		right--
		if top <= bottom {
			for i:=right; i>=left; i-- {
				res[bottom][i] = val
				val++
			}
			bottom--
		}
		if left <= right {
			for i:=bottom; i>=top; i-- {
				res[i][left] = val
				val++
			}
			left++
		}
	}
	return res
}