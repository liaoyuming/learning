package leetcode

/**
		不同路径
		https://leetcode-cn.com/problems/unique-paths/
 */

func main() {
	println(uniquePaths(1, 1))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for k,_ := range dp {
		dp[k] = make([]int, n)
	}
	for i:=0; i<m; i++ {
		dp[i][0] = 1
	}
	for i:=0; i<n; i++ {
		dp[0][i] = 1
	}
	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}