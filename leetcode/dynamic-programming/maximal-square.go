package main

func maximalSquare(matrix [][]byte) int {
	max := 0
	dp := make([][]int, len(matrix))
	for i:=0; i<len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix[i]); j++ {
			if i==0 || j==0 {
				dp[i][j] = int(matrix[i][j] - '0')
			} else {
				if matrix[i][j] == '1' {
					dp[i][j] = int(min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1)
				} else {
					dp[i][j] = 0
				}
			}
			if max < dp[i][j] {
				max = dp[i][j]
			}
		}
	}
	return max * max
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	println(maximalSquare([][]byte{
		{'1', '1'},
		{'1', '1'},
		//{'1', '0', '1', '1', '1'},
		//{'1', '1', '1', '1', '1'},
		//{'1', '0', '0', '1', '0'},
	}))
}