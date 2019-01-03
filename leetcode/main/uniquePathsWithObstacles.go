package leetcode

/**
		不同路径 II
		https://leetcode-cn.com/problems/unique-paths-ii/
 */

func main() {
	println(uniquePathsWithObstacles([][]int{
		{1,0},
	}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for k,_ := range dp {
		dp[k] = make([]int, n)
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}