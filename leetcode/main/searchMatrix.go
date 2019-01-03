package leetcode

/**
		搜索二维矩阵
		https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 */

func main() {
	var matrix [][]int
	//matrix = append(matrix, []int{1})
	matrix = append(matrix, []int{10,11,16,20})
	matrix = append(matrix, []int{23,30,34,50})
	matrix = append(matrix, []int{39, 40, 44, 50})
	println(searchMatrix(matrix, 29))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	start := 0
	end := m*n-1
	mid := 0
	for start <= end {
		mid = start + (end - start) / 2
		if matrix[mid/n][mid%n] < target {
			start = mid+1
		} else if matrix[mid/n][mid%n] > target {
			end = mid-1
		} else {
			return true
		}
	}
	return false
}