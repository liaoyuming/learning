package main

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := make([][]int, len(grid))
	for k,_ := range dp {
		dp[k] = make([]int, len(grid[k]))
	}
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[i]); j++ {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a int, b int) int {
	if a<b {
		return a
	}
	return b
}