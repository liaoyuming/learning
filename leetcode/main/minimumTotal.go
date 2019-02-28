package main

/**
		三角形最小路径和
		https://leetcode-cn.com/problems/triangle/description/
 */

func main() {
	println(minimumTotal([][]int{
		{2},
		//{3,4},
		//{6,5,7},
		//{4,1,8,3},
	}))
}


func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	dp := make([]int, length)
	dp[0] = triangle[0][0]
	a := 0
	b := 0
	minTotal := dp[0]
	for i:=1; i<length; i++ {
		for j:=0; j<len(triangle[i]); j++ {
			a = b
			b = dp[j]
			if j == 0 {
				dp[j] = b + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[j] = a + triangle[i][j]
			} else {
				dp[j] = min(a, b) + triangle[i][j]
			}
			if i == length-1 {
				if j == 0 {
					minTotal = dp[0]
				} else {
					minTotal = min(minTotal, dp[j])
				}
			}
		}
	}
	return minTotal
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}