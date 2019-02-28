package main

/**
		不同的二叉搜索树
		https://leetcode-cn.com/problems/unique-binary-search-trees/
 */

func main() {
	println(numTrees(1000))
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	// 动态规划 方程： 当前个数等于所有可能性的左右子树的个数之和
	for i:=2; i<=n; i++ {
		for j:=0; j<i; j++ {
			dp[i] += dp[j]*dp[i-j-1]
		}
	}
	return dp[n]
}